// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/cloyop/tiniest/internal/templates/shared"

func TermsOfService() templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"flex flex-col justify-center items-center w-full min-h-screen bg-bluer\"><a href=\"/\" class=\"absolute top-4 left-4\">⬅</a><div class=\"flex flex-col gap-4 w-10/12 md:w-8/12 font-medium shadow-b-outer2 bg-blue-dk py-4 px-2 text-aqua-light-3\"><section class=\"flex flex-col gap-2\"><h1 class=\"text-3xl text-white font-bold\">Terms of Service</h1><p>This page outlines the terms and conditions for using Short URL's URL\r shortening service. By accessing and using our website, you agree to\r comply with these guidelines.\r</p></section><section><h2 class=\"text-2xl text-white font-semibold\">Prohibited Use</h2><div class=\"h-2\"></div><p>Our service is available to all users, however, in order to ensure the\r safety and security of our users, certain types of content and usage\r are prohibited.\r</p></section><section><p class=\"text-xl font-semibold text-white\">It is prohibited to create links that redirect to:\r</p><ul class=\"flex flex-col gap-1 pl-2\"><li>• Abusive or suspicious content</li><li>• Pornographic, sexual, violent or prejudiced content</li><li>• Content related to drugs, weapons or illegal substances</li><li>• Child or child-related content</li><li>• Double redirection or pages without content</li><li>• Pages without SSL</li></ul></section><section class=\"flex flex-col gap-2\"><p>All URLs will be analyzed. Any shortened URL that leads to a page that\r violates this conditions may be deleted without notice. Shortened URLs\r that do not receive at least 1 click per month are also removed.\r</p></section></div></body>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = shared.Page("Tiniest • Terms Of Service").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
