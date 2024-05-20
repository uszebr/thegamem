package modelfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uszebr/thegamem/play/model/implementation/modalwaysgreen"
	"github.com/uszebr/thegamem/play/model/implementation/modalwaysred"
	"github.com/uszebr/thegamem/play/model/implementation/modblindrevenge"
	"github.com/uszebr/thegamem/play/model/implementation/modcopystartgreen"
	"github.com/uszebr/thegamem/play/model/implementation/modcopystartred"
	"github.com/uszebr/thegamem/play/model/implementation/modrandom"
)

func TestGetModelFactory(t *testing.T) {
	factory := GetModelFactory()
	anotherFactory := GetModelFactory()
	if factory == nil || anotherFactory == nil {
		t.Errorf("Model Factory is nill")
	}
	if factory != anotherFactory {
		t.Errorf("Different factories from singleton factory: %v factory1: %v", factory, anotherFactory)
	}
}

func TestMustCreateModel(t *testing.T) {

	factory := GetModelFactory()

	tests := []struct {
		modelName       string
		expectedType    ModelCreatorI
		expectedToPanic bool
		iconPath        string
	}{
		{"alwaysgreen", modalwaysgreen.ModAlwaysGreen{}, false, "/static/image/icon/rabbit.svg"},
		{"alwaysred", modalwaysred.ModAlwaysRed{}, false, "/static/image/icon/crocodile.svg"},
		{"blindrevenge", modblindrevenge.ModBlindRevenge{}, false, "/static/image/icon/pig.svg"},
		{"copystartgreen", modcopystartgreen.ModCopyStrartGreen{}, false, "/static/image/icon/whale.svg"},
		{"copystartred", modcopystartred.ModCopyStrartRed{}, false, "/static/image/icon/fox.svg"},
		{"random", modrandom.ModRandom{}, false, "/static/image/icon/jellyfish.svg"},
		{"nonexistent", nil, true, ""},
		//to fail {"random", modrandom.ModRandom{}, false, "/static/image/icon/fakeicon.svg"},
	}

	for _, tt := range tests {
		if tt.expectedToPanic {
			assert.Panics(t, func() {
				factory.MustCreateModel(tt.modelName)
			}, "Expected panic for modelName: %s", tt.modelName)
		} else {
			model := factory.MustCreateModel(tt.modelName)
			assert.IsType(t, tt.expectedType.GetModel(), model, "Expected model type for modelName: %s", tt.modelName)
			assert.Equalf(t, model.GetIcon(), tt.iconPath, "Expected Icon path: %s actual: ", tt.iconPath, model.GetIcon())
		}
	}
}
