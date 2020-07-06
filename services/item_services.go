package services

var (
	ItemService itemServiceInterface = &itemService{}
)

type itemService struct{}

type itemServiceInterface interface {
	GetItem()
	SearchItem()
}

func (i *itemService) GetItem() {

}

func (i *itemService) SearchItem() {

}
