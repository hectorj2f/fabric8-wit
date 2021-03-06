package design

import (
	d "github.com/goadesign/goa/design"
	a "github.com/goadesign/goa/design/apidsl"
)

var searchWorkItemList = JSONList(
	"SearchWorkItem", "Holds the paginated response to a search request",
	workItem2,
	pagingLinks,
	meta)

var _ = a.Resource("search", func() {
	a.BasePath("/search")

	a.Action("show", func() {
		a.Routing(
			a.GET(""),
		)
		a.Description("Search by ID, URL, full text capability")
		a.Params(func() {
			a.Param("q", d.String,
				`Following are valid input for search query
				1) "id:100" :- Look for work item hainvg id 100
				2) "url:http://demo.almighty.io/details/500" :- Search on WI having id 500 and check 
					if this URL is mentioned in searchable columns of work item
				3) "simple keywords separated by space" :- Search in Work Items based on these keywords.`)
			a.Param("page[offset]", d.String, "Paging start position") // #428
			a.Param("page[limit]", d.Integer, "Paging size")
			a.Required("q")
		})
		a.Response(d.OK, func() {
			a.Media(searchWorkItemList)
		})
		a.Response(d.BadRequest, JSONAPIErrors)
		a.Response(d.InternalServerError, JSONAPIErrors)
	})
})
