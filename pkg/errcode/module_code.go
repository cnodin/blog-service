package errcode

var (
	ErrorGetTagListFail	=	NewError(20010001, "Get tags list failture")
	ErrorCreateTagFail	=	NewError(20010002, "Create tag failture")
	ErrorUpdateTagFail	=	NewError(20010003, "Update tag failture")
	ErrorDeleteTagFail	=	NewError(20010004, "Delete tag failture")
	ErrorCountTagFail	=	NewError(20010005, "Count tag failture")

	ErrorGetArticleFail    = NewError(20020001, "Get article failture")
	ErrorGetArticlesFail   = NewError(20020002, "Get mutil-articles failture")
	ErrorCreateArticleFail = NewError(20020003, "Create article failture")
	ErrorUpdateArticleFail = NewError(20020004, "Update article failture")
	ErrorDeleteArticleFail = NewError(20020005, "Delete article failture")

	ErrorUploadFileFail = NewError(20030001, "Upload file failture")
)

