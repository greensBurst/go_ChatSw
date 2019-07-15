package process
import (
	"fmt"
	"go_ChatSw/public"
	"encoding/json"
)

func outputGroupMes(mes *public.Message) {

	var smsMes public.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal error:",err)
		return
	}

	info := fmt.Sprintf("用户id:\t%s 对大家说:\t%s",
		smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
