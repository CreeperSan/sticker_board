package AccountPackage

var _instance *AccountInterface

func InstallOperator(instance *AccountInterface){
	_instance = instance
}

func UninstallOperator(){
	_instance = nil
}

func GetOperator() *AccountInterface {
	if _instance != nil {
		return _instance
	}

	panic("Account Package has not been install yet!")
}

