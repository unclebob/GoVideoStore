package GoVideoStore

type RentalTypeFactoryImpl struct{}

func (r RentalTypeFactoryImpl) make(typeName string) RentalType {
	switch typeName {
	case "childrens":
		return ChildrensRental{}
	case "new release":
		return NewReleaseRental{}
	case "regular":
		return RegularRental{}
	}
	return nil
}
