// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package navmenuview

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/uszebr/thegamem/internal/entity"

func Show() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"navbar navbar-expand-lg bg-body-tertiary\"><div class=\"container-fluid\"><a class=\"navbar-brand\" title=\"Game Theory in Memory Calculation(noDB playground)\" href=\"/\">THEGAMEM</a> <button class=\"navbar-toggler\" type=\"button\" data-bs-toggle=\"collapse\" data-bs-target=\"#navbarNavDropdown\" aria-controls=\"navbarNavDropdown\" aria-expanded=\"false\" aria-label=\"Toggle navigation\"><span class=\"navbar-toggler-icon\"></span></button><div class=\"collapse navbar-collapse\" id=\"navbarNavDropdown\"><ul class=\"navbar-nav\"><li class=\"nav-item\"><a class=\"nav-link active\" aria-current=\"page\" href=\"/\">Home</a></li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if _, ok := ctx.Value("user").(entity.UserAuth); ok {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item\"><a class=\"nav-link\" href=\"#\" hx-post=\"/logout\" hx-swap=\"none\" hx-trigger=\"click\">Logout</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/newgame\">New Game</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/game\">Game</a></li><li class=\"nav-item\"><a class=\"nav-link\" href=\"/statistics\">Statistics</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item\"><a class=\"nav-link\" href=\"/login\">Login</a></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"nav-item\"><a class=\"nav-link\" href=\"#\">About</a></li></ul></div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
