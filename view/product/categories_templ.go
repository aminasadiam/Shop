// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package product

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/aminasadiam/Shop/view"

func Categories() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-6\"><div class=\"flex justify-between mb-6 col-span-full\"><h1 class=\"text-3xl font-bold\">Product Categories</h1></div><div class=\"bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300\"><div class=\"relative\"><img src=\"https://via.placeholder.com/300\" alt=\"Electronics\" class=\"w-full h-48 object-cover\"><div class=\"absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center\"><h3 class=\"text-2xl font-bold text-white\">Electronics</h3></div></div><div class=\"p-4\"><p class=\"text-gray-600 text-sm\">Browse our collection of latest electronics and gadgets</p><a href=\"/products?category=electronics\" class=\"mt-4 inline-block text-blue-500 hover:text-blue-600\">View Products →</a></div></div><div class=\"bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300\"><div class=\"relative\"><img src=\"https://via.placeholder.com/300\" alt=\"Clothing\" class=\"w-full h-48 object-cover\"><div class=\"absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center\"><h3 class=\"text-2xl font-bold text-white\">Clothing</h3></div></div><div class=\"p-4\"><p class=\"text-gray-600 text-sm\">Discover trendy and fashionable clothing items</p><a href=\"/products?category=clothing\" class=\"mt-4 inline-block text-blue-500 hover:text-blue-600\">View Products →</a></div></div><div class=\"bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300\"><div class=\"relative\"><img src=\"https://via.placeholder.com/300\" alt=\"Home &amp; Living\" class=\"w-full h-48 object-cover\"><div class=\"absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center\"><h3 class=\"text-2xl font-bold text-white\">Home & Living</h3></div></div><div class=\"p-4\"><p class=\"text-gray-600 text-sm\">Everything you need to make your home beautiful</p><a href=\"/products?category=home-living\" class=\"mt-4 inline-block text-blue-500 hover:text-blue-600\">View Products →</a></div></div><div class=\"bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300\"><div class=\"relative\"><img src=\"https://via.placeholder.com/300\" alt=\"Books\" class=\"w-full h-48 object-cover\"><div class=\"absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center\"><h3 class=\"text-2xl font-bold text-white\">Books</h3></div></div><div class=\"p-4\"><p class=\"text-gray-600 text-sm\">Explore our vast collection of books across genres</p><a href=\"/products?category=books\" class=\"mt-4 inline-block text-blue-500 hover:text-blue-600\">View Products →</a></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = view.Base("Categories").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
