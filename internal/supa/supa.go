package supa

import (
	"context"
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/nedpals/supabase-go"
	"github.com/uszebr/thegamem/internal/entity"
)

// Unofficial Go Supabase client
// Has Auth
// Has strange parsing model
// Has return of inserted row data!! might be helpful
var (
	client *supabase.Client
	once   sync.Once
)

type SupabaseAuthWrapper struct {
	originSupa *supabase.Auth
}

func getClient() *supabase.Client {
	once.Do(func() {
		url := os.Getenv("SUPABASE_URL")
		key := os.Getenv("SUPABASE_PUBLIC_KEY")
		if url == "" || key == "" {
			log.Fatalf("Supa key or url is missing")
		}
		supabase := supabase.CreateClient(url, key)
		client = supabase
	})
	return client
}

func GetSupaAuth() *SupabaseAuthWrapper {
	supaAuth := getClient().Auth
	slog.Debug("Supabase DB is set")
	return &SupabaseAuthWrapper{originSupa: supaAuth}
}

func (supabaseAuthWrapper SupabaseAuthWrapper) RefreshUser(ctx context.Context, userToken string, refreshToken string) (*entity.AuthDetails, error) {
	supaDetails, err := supabaseAuthWrapper.originSupa.RefreshUser(ctx, userToken, refreshToken)
	return &entity.AuthDetails{AccessToken: supaDetails.AccessToken, ExpiresIn: supaDetails.ExpiresIn, RefreshToken: supaDetails.RefreshToken}, err
}

func (supabaseAuthWrapper SupabaseAuthWrapper) SignIn(ctx context.Context, userLogin string, userPass string) (*entity.AuthDetails, error) {
	supaDetails, err := supabaseAuthWrapper.originSupa.SignIn(context.Background(), supabase.UserCredentials{Email: userLogin, Password: userPass})
	if err != nil {
		return &entity.AuthDetails{}, err
	}
	return &entity.AuthDetails{AccessToken: supaDetails.AccessToken, ExpiresIn: supaDetails.ExpiresIn, RefreshToken: supaDetails.RefreshToken}, nil
}

func (supabaseAuthWrapper SupabaseAuthWrapper) SignOut(ctx context.Context, userTocken string) error {
	return supabaseAuthWrapper.originSupa.SignOut(context.Background(), userTocken)
}
