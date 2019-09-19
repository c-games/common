package err

// app、service
const Code_Success int = 0

const (
	Code_No_Such_Command int = 3001

	// user or admin
	Code_User_Not_exist int = 3000
	Code_User_Login_Failed int = 3000
	Code_User_Id_Not_Exist int = 3000
	Code_User_Already_exist int = 3000
	Code_Wrong_Token int = 3000
	Code_Permission_Denied int = 3000
	Code_Wrong_Account_Type int = 3000
	Code_Wrong_Request_Data int = 3000
	Code_Not_Found int = 3000
	Code_Unmarshal_Failed int = 3000
	Code_Marshal_Failed int = 3000
	Code_No_Token int = 3000
	Code_Multiple_Query_Failed int = 3000
)

// Data
const (
	Code_Unexpected_Data int = 3000
	Code_Not_Homo_Data int = 3000
)

// DB
const (
	Code_DB_Error int = 4000
)

// Calculate
const (

)