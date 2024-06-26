// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package helpers

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ViewsSvg() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"pt-0.5\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"18px\" height=\"auto\" fill=\"none\" viewBox=\"0 0 256 107\"><path stroke=\"#D1FFF0\" stroke-width=\"8\" d=\"m3 55.666 25.044 16.59a182.312 182.312 0 0 0 205.172-2.587l20.022-14.003\"></path> <path stroke=\"#D1FFF0\" stroke-width=\"8\" d=\"m3 59.27 27.477-21.444a159.81 159.81 0 0 1 199.944 2.645l22.817 18.8\"></path> <path stroke=\"#D1FFF0\" stroke-width=\"10\" d=\"M181.533 56.249c0 6.195-4.654 12.895-14.595 18.336-9.756 5.34-23.574 8.798-39.111 8.798-15.537 0-29.356-3.458-39.112-8.798-9.94-5.441-14.595-12.141-14.595-18.336 0-6.195 4.654-12.895 14.595-18.336 9.756-5.34 23.575-8.798 39.112-8.798 15.537 0 29.355 3.458 39.111 8.798 9.941 5.441 14.595 12.14 14.595 18.336Z\"></path></svg></span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func GitHubSVG() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg class=\"bg-black rounded-full\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\"><path d=\"M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z\" fill=\"#d6d3d1\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func GoogleSVG() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" shape-rendering=\"geometricPrecision\" text-rendering=\"geometricPrecision\" image-rendering=\"optimizeQuality\" fill-rule=\"evenodd\" clip-rule=\"evenodd\" viewBox=\"0 0 640 640\"><path d=\"M326.331 274.255v109.761h181.49c-7.37 47.115-54.886 138.002-181.49 138.002-109.242 0-198.369-90.485-198.369-202.006 0-111.509 89.127-201.995 198.369-201.995 62.127 0 103.761 26.516 127.525 49.359l86.883-83.635C484.99 31.512 412.741-.012 326.378-.012 149.494-.012 6.366 143.116 6.366 320c0 176.884 143.128 320.012 320.012 320.012 184.644 0 307.256-129.876 307.256-312.653 0-21-2.244-36.993-5.008-52.997l-302.248-.13-.047.024z\" fill=\"#d6d3d1\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func PadUnlock() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16px\" height=\"16px\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M8.18185 10.7027H6.00004C5.44776 10.7027 5.00005 11.1485 5.00004 11.7008C5.00001 13.3483 4.99996 16.6772 5.00006 18.9189C5.00018 21.4317 8.88614 22 12 22C15.1139 22 18.9999 21.4317 18.9999 18.9189C19 16.6773 19 13.3483 18.9999 11.7008C18.9999 11.1485 18.5522 10.7027 17.9999 10.7027H15.8181H8.18185ZM8.18185 10.7027C8.18185 10.7027 8.18189 8.13513 8.18185 6.59459C8.18181 4.74571 9.70882 3 12 3C14.2912 3 15.8181 4.74571 15.8181 6.59459\" stroke=\"#d1fff0\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M13 16.6181V18C13 18.5523 12.5523 19 12 19C11.4477 19 11 18.5523 11 18V16.6181C10.6931 16.3434 10.5 15.9442 10.5 15.5C10.5 14.6716 11.1716 14 12 14C12.8284 14 13.5 14.6716 13.5 15.5C13.5 15.9442 13.3069 16.3434 13 16.6181Z\" fill=\"#d1fff0\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func PadLock() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16px\" height=\"16px\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M8.1819 10.7027H6.00008C5.44781 10.7027 5.0001 11.1485 5.00009 11.7008C5.00005 13.3483 5 16.6772 5.00011 18.9189C5.00023 21.4317 8.88618 22 12 22C15.1139 22 19 21.4317 19 18.9189C19 16.6773 19 13.3483 19 11.7008C19 11.1485 18.5523 10.7027 18 10.7027H15.8182M8.1819 10.7027C8.1819 10.7027 8.18193 8.13514 8.1819 6.59459C8.18186 4.74571 9.70887 3 12 3C14.2912 3 15.8182 4.74571 15.8182 6.59459C15.8182 8.13514 15.8182 10.7027 15.8182 10.7027M8.1819 10.7027H15.8182\" stroke=\"#d1fff0\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M13 16.6181V18C13 18.5523 12.5523 19 12 19C11.4477 19 11 18.5523 11 18V16.6181C10.6931 16.3434 10.5 15.9442 10.5 15.5C10.5 14.6716 11.1716 14 12 14C12.8284 14 13.5 14.6716 13.5 15.5C13.5 15.9442 13.3069 16.3434 13 16.6181Z\" fill=\"#d1fff0\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func AnalyticSVG() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"14px\" height=\"14px\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><g fill=\"#d1fff0\" fill-rule=\"evenodd\"><path d=\"M24 23a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V1a1 1 0 1 1 2 0v21h21a1 1 0 0 1 1 1z\"></path> <path d=\"M23.6 1.8l-2.62 1.969a.97.97 0 0 1 .02.24 2.006 2.006 0 0 1-2 2A.569.569 0 0 1 18.871 6l-1.04 1.559-1.04 1.559c.14.273.211.576.209.882a1.994 1.994 0 0 1-1.88 1.989l-1.45 2.9A2 2 0 1 1 10 16v-.03l-2.03-1.218c-.177.1-.37.172-.57.21L5.95 19.32A1 1 0 0 1 5 20a1.252 1.252 0 0 1-.32-.05 1.01 1.01 0 0 1-.63-1.269l1.45-4.358A2 2 0 1 1 9 13v.03l2.03 1.219c.26-.147.552-.23.85-.24l1.45-2.9A2.006 2.006 0 0 1 15 8.005a.569.569 0 0 1 .13.01l1.04-1.559L17.211 4.9a2 2 0 0 1 2.569-2.732L22.4.2a1 1 0 1 1 1.2 1.6z\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CopySVG(cp string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg data-cp=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(cp)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/helpers/svg.templ`, Line: 89, Col: 14}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"18px\" height=\"18px\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path data-cp=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(cp)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/helpers/svg.templ`, Line: 97, Col: 15}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M19.5 16.5L19.5 4.5L18.75 3.75H9L8.25 4.5L8.25 7.5L5.25 7.5L4.5 8.25V20.25L5.25 21H15L15.75 20.25V17.25H18.75L19.5 16.5ZM15.75 15.75L15.75 8.25L15 7.5L9.75 7.5V5.25L18 5.25V15.75H15.75ZM6 9L14.25 9L14.25 19.5L6 19.5L6 9Z\" fill=\"#d1fff0\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func DeleteSVG() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg fill=\"#d1fff0\" width=\"16px\" height=\"16px\" viewBox=\"-3.5 0 19 19\" xmlns=\"http://www.w3.org/2000/svg\" class=\"cf-icon-svg\"><path d=\"M11.383 13.644A1.03 1.03 0 0 1 9.928 15.1L6 11.172 2.072 15.1a1.03 1.03 0 1 1-1.455-1.456l3.928-3.928L.617 5.79a1.03 1.03 0 1 1 1.455-1.456L6 8.261l3.928-3.928a1.03 1.03 0 0 1 1.455 1.456L7.455 9.716z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func DownloadSVG() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var11 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var11 == nil {
			templ_7745c5c3_Var11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"14px\" height=\"14px\" viewBox=\"0 0 192 192\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\"><g stroke=\"#d1fff0\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"20\" clip-path=\"url(#a)\"><path d=\"m30 104 66 66 66-66m-66 52V22\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Loader_sm(id string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var12 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var12 == nil {
			templ_7745c5c3_Var12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"htmx-indicator\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(id)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/helpers/svg.templ`, Line: 133, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><svg aria-hidden=\"true\" class=\"w-4 h-auto text-gray-200 animate-spin dark:text-gray-600 fill-blue-600\" viewBox=\"0 0 100 101\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z\" fill=\"currentColor\"></path> <path d=\"M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z\" fill=\"currentFill\"></path></svg></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Loader(id string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var14 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var14 == nil {
			templ_7745c5c3_Var14 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"htmx-indicator\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(id)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/helpers/svg.templ`, Line: 154, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><svg aria-hidden=\"true\" class=\"w-6 h-6 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600\" viewBox=\"0 0 100 101\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z\" fill=\"currentColor\"></path> <path d=\"M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z\" fill=\"currentFill\"></path></svg></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ProtectCardSVG() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var16 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var16 == nil {
			templ_7745c5c3_Var16 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"36\" height=\"32\" fill=\"none\" viewBox=\"0 0 45 45\"><path stroke=\"#7BFED3\" stroke-width=\"2\" d=\"M22.5 1H44v31.022L22.5 43.859 1 32.022V1h21.5Z\"></path> <path stroke=\"#7BFED3\" stroke-width=\"2\" d=\"M22.5 6.4h14.75v21.898L22.5 36.651 7.75 28.298V6.4H22.5Z\"></path> <path stroke=\"#7BFED3\" stroke-width=\"2\" d=\"M22.5 3.7h18.125v26.46L22.5 40.255 4.375 30.16V3.7H22.5Z\"></path> <path stroke=\"#7BFED3\" stroke-width=\"2\" d=\"M22.5 9.1h11.375v17.34L22.5 33.044 11.125 26.44V9.1H22.5Z\"></path> <path stroke=\"#7BFED3\" stroke-width=\"2\" d=\"M22.5 11.8h8v12.786l-8 4.845-8-4.845V11.8h8Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func QRCardSVG() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var17 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var17 == nil {
			templ_7745c5c3_Var17 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"36\" height=\"32\" fill=\"none\" viewBox=\"0 0 40 40\"><g clip-path=\"url(#a)\"><path fill=\"#000A11\" d=\"M0 0h40v40H0z\"></path> <path stroke=\"#1BF8B0\" d=\"M39 8V3a2 2 0 0 0-2-2h-5M9 1H3a2 2 0 0 0-2 2v5m0 24v4a3 3 0 0 0 3 3h3m32-7v5a2 2 0 0 1-2 2h-4\"></path> <rect width=\"9\" height=\"9\" x=\"15.5\" y=\"15.5\" stroke=\"#1BF8B0\" rx=\"2.5\"></rect> <circle cx=\"20\" cy=\"20\" r=\".5\" fill=\"#020F19\" stroke=\"#1BF8B0\"></circle></g> <defs><clipPath id=\"a\"><path fill=\"#fff\" d=\"M0 0h40v40H0z\"></path></clipPath></defs></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CopyCardSVG() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var18 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var18 == nil {
			templ_7745c5c3_Var18 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"36\" height=\"32\" fill=\"none\" viewBox=\"0 0 52 36\"><path stroke=\"url(#a)\" stroke-linecap=\"round\" d=\"m10.487 18.636-9.4 16.239\"></path> <path stroke=\"url(#b)\" stroke-linecap=\"round\" d=\"M1 35.102h41.045\"></path> <path stroke=\"url(#c)\" stroke-linecap=\"round\" d=\"m51 18.613-8.713 16.354\"></path> <path stroke=\"url(#d)\" stroke-linecap=\"round\" d=\"M10.958 18.613h39.854\"></path> <path stroke=\"url(#e)\" stroke-linecap=\"round\" d=\"M1 34.872V10.868\"></path> <path stroke=\"url(#f)\" stroke-linecap=\"round\" d=\"M4.73 28.231V7.371\"></path> <path stroke=\"url(#g)\" stroke-linecap=\"round\" d=\"M8.09 22.61V5.247\"></path> <path stroke=\"url(#h)\" stroke-linecap=\"round\" d=\"M1 10.743h3.248\"></path> <path stroke=\"url(#i)\" stroke-linecap=\"round\" d=\"M4.73 7.37h2.8\"></path> <path stroke=\"url(#j)\" stroke-linecap=\"round\" d=\"M47.083 13.241v4.622\"></path> <path stroke=\"url(#k)\" stroke-linecap=\"round\" d=\"M39.247 13.241h7.836\"></path> <path stroke=\"url(#l)\" stroke-linecap=\"round\" d=\"M15.615 4.248h5.599\"></path> <path stroke=\"url(#m)\" stroke-linecap=\"round\" d=\"M15.615 7.496h8.398\"></path> <path stroke=\"url(#n)\" stroke-linecap=\"round\" d=\"M15.615 10.743h12.131\"></path> <path stroke=\"url(#o)\" stroke-linecap=\"round\" d=\"M15.615 13.991h12.131\"></path> <path stroke=\"url(#p)\" stroke-linecap=\"round\" d=\"M12.566 18.613V1\"></path> <path stroke=\"url(#q)\" stroke-linecap=\"round\" d=\"M38.872 1H12.566\"></path> <path stroke=\"url(#r)\" stroke-linecap=\"round\" d=\"M38.872 18.613V1\"></path> <path stroke=\"url(#s)\" stroke-linecap=\"round\" d=\"M8.308 5.06h4.06\"></path> <defs><linearGradient id=\"a\" x1=\"5.788\" x2=\"4.954\" y1=\"26.756\" y2=\"26.273\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"b\" x1=\"21.522\" x2=\"21.522\" y1=\"35.102\" y2=\"36.102\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"c\" x1=\"46.643\" x2=\"45.789\" y1=\"26.79\" y2=\"26.335\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"d\" x1=\"30.885\" x2=\"30.885\" y1=\"18.613\" y2=\"19.613\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"e\" x1=\"1\" x2=\"2\" y1=\"22.87\" y2=\"22.87\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"f\" x1=\"4.731\" x2=\"5.731\" y1=\"17.801\" y2=\"17.801\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"g\" x1=\"8.089\" x2=\"9.089\" y1=\"13.928\" y2=\"13.928\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"h\" x1=\"2.624\" x2=\"2.624\" y1=\"10.743\" y2=\"11.743\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"i\" x1=\"6.13\" x2=\"6.13\" y1=\"7.371\" y2=\"8.371\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"j\" x1=\"47.083\" x2=\"46.083\" y1=\"15.552\" y2=\"15.552\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"k\" x1=\"43.165\" x2=\"43.165\" y1=\"13.241\" y2=\"14.241\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"l\" x1=\"18.415\" x2=\"18.415\" y1=\"4.248\" y2=\"5.247\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"m\" x1=\"19.814\" x2=\"19.814\" y1=\"7.496\" y2=\"8.495\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"n\" x1=\"21.681\" x2=\"21.681\" y1=\"10.743\" y2=\"11.743\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"o\" x1=\"21.681\" x2=\"21.681\" y1=\"13.991\" y2=\"14.991\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"p\" x1=\"12.566\" x2=\"13.566\" y1=\"9.806\" y2=\"9.806\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"q\" x1=\"25.719\" x2=\"25.719\" y1=\"1\" y2=\"0\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"r\" x1=\"38.872\" x2=\"39.872\" y1=\"9.806\" y2=\"9.806\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient> <linearGradient id=\"s\" x1=\"10.338\" x2=\"10.338\" y1=\"5.06\" y2=\"6.06\" gradientUnits=\"userSpaceOnUse\"><stop stop-color=\"#D1FFF0\"></stop> <stop offset=\"1\" stop-color=\"#FBFF35\"></stop></linearGradient></defs></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func TrackingSVG() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var19 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var19 == nil {
			templ_7745c5c3_Var19 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"36px\" height=\"36px\" viewBox=\"0 -0.5 17 17\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" class=\"si-glyph si-glyph-foot-sign\"><g transform=\"translate(1.000000, 0.000000)\" fill=\"#8cffd9\"><path d=\"M4.428,13.572 L0.629,12.142 L0.145,13.315 C0.145,13.315 -0.318,15.213 1.342,15.838 C3.004,16.465 3.961,14.751 3.961,14.751 L4.428,13.572 L4.428,13.572 Z\"></path> <path d=\"M7.207,3.193 C5.565,2.534 3.26,3.979 2.463,5.8 C2.135,6.55 1.986,7.359 1.862,8.157 C1.803,8.538 1.761,8.929 1.686,9.309 C1.59,9.786 1.447,10.245 1.305,10.708 C1.108,11.351 1.325,11.459 1.924,11.569 L4.022,12.361 C4.236,12.463 4.654,12.72 4.869,12.48 C5.059,12.265 5.021,11.873 5.148,11.618 C5.312,11.287 5.496,10.95 5.699,10.638 C6.148,9.94 7,9.43 7.577,8.828 C8.292,8.08 8.687,7.33 8.905,6.338 C9.195,5.017 8.528,3.722 7.207,3.193 L7.207,3.193 Z\"></path> <g transform=\"translate(8.000000, 0.000000)\"><path d=\"M0.977,9.289 L4.632,10.732 C4.632,10.732 3.878,13.685 1.646,12.826 C-0.586,11.965 0.977,9.289 0.977,9.289 L0.977,9.289 Z\"></path> <path d=\"M6.19,0.217 C7.75,0.797 8.378,3.255 7.721,5.024 C7.45,5.751 7.018,6.403 6.575,7.038 C6.363,7.34 6.133,7.636 5.932,7.949 C5.685,8.339 5.479,8.75 5.271,9.16 C4.98,9.73 4.759,9.665 4.275,9.366 L2.31,8.593 C2.097,8.529 1.641,8.441 1.653,8.142 C1.664,7.872 1.949,7.622 2.031,7.368 C2.137,7.035 2.234,6.683 2.3,6.34 C2.452,5.572 2.204,4.679 2.208,3.899 C2.208,2.93 2.435,2.159 2.94,1.334 C3.617,0.228 4.932,-0.248 6.19,0.217 L6.19,0.217 Z\"></path></g></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func TrackingSVG_SM() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var20 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var20 == nil {
			templ_7745c5c3_Var20 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"14px\" height=\"14px\" viewBox=\"0 -0.5 17 17\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" class=\"si-glyph si-glyph-foot-sign\"><g transform=\"translate(1.000000, 0.000000)\" fill=\"#d1fff0\"><path d=\"M4.428,13.572 L0.629,12.142 L0.145,13.315 C0.145,13.315 -0.318,15.213 1.342,15.838 C3.004,16.465 3.961,14.751 3.961,14.751 L4.428,13.572 L4.428,13.572 Z\"></path> <path d=\"M7.207,3.193 C5.565,2.534 3.26,3.979 2.463,5.8 C2.135,6.55 1.986,7.359 1.862,8.157 C1.803,8.538 1.761,8.929 1.686,9.309 C1.59,9.786 1.447,10.245 1.305,10.708 C1.108,11.351 1.325,11.459 1.924,11.569 L4.022,12.361 C4.236,12.463 4.654,12.72 4.869,12.48 C5.059,12.265 5.021,11.873 5.148,11.618 C5.312,11.287 5.496,10.95 5.699,10.638 C6.148,9.94 7,9.43 7.577,8.828 C8.292,8.08 8.687,7.33 8.905,6.338 C9.195,5.017 8.528,3.722 7.207,3.193 L7.207,3.193 Z\"></path> <g transform=\"translate(8.000000, 0.000000)\"><path d=\"M0.977,9.289 L4.632,10.732 C4.632,10.732 3.878,13.685 1.646,12.826 C-0.586,11.965 0.977,9.289 0.977,9.289 L0.977,9.289 Z\"></path> <path d=\"M6.19,0.217 C7.75,0.797 8.378,3.255 7.721,5.024 C7.45,5.751 7.018,6.403 6.575,7.038 C6.363,7.34 6.133,7.636 5.932,7.949 C5.685,8.339 5.479,8.75 5.271,9.16 C4.98,9.73 4.759,9.665 4.275,9.366 L2.31,8.593 C2.097,8.529 1.641,8.441 1.653,8.142 C1.664,7.872 1.949,7.622 2.031,7.368 C2.137,7.035 2.234,6.683 2.3,6.34 C2.452,5.572 2.204,4.679 2.208,3.899 C2.208,2.93 2.435,2.159 2.94,1.334 C3.617,0.228 4.932,-0.248 6.19,0.217 L6.19,0.217 Z\"></path></g></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
