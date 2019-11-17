package vm

// BaseViewModel struct
type BaseViewModel struct {
	Title       string
	CurrentUser string
}

// SetTitle func
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}

// SetCurrentUser func
func (v *BaseViewModel) SetCurrentUser(username string) {
	v.CurrentUser = username
}

type BasePageViewModel struct {
	PrevPage    int
	NextPage    int
	Total       int //总页数
	CurrentPage int
	Limit       int //每页显示项目数
}

func (v *BasePageViewModel) SetPrevAndNextPage() {
	if v.CurrentPage > 1 {
		v.PrevPage = v.CurrentPage - 1
	}
	if (v.Total-1)/v.Limit >= v.CurrentPage {
		v.NextPage = v.CurrentPage + 1
	}
}

func (v *BasePageViewModel) SetBasePageViewModel(total, page, limit int) {
	v.Total = total
	v.CurrentPage = page
	v.Limit = limit
	v.SetPrevAndNextPage()
}
