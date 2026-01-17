package main

import (
	"context"
	"os"

	// "time"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func CompleteGenerate(ctx context.Context, model *ark.ChatModel, input []*schema.Message) {
	resp, err := model.Generate(ctx, input)
	if err != nil {
		panic(err)
	}
	print(resp.Content)
}

func StreamGenerate(ctx context.Context, model *ark.ChatModel, input []*schema.Message) {
	reader, err := model.Stream(ctx, input)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	for {
		msg, err := reader.Recv()
		if err != nil {
			panic(err)
		}
		print(msg.Content)
		// time.Sleep(time.Millisecond * 100)
	}
}
func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	model, err := ark.NewChatModel(ctx,
		&ark.ChatModelConfig{
			APIKey: os.Getenv("ARK_API_KEY"), // 修正环境变量名称
			Model:  os.Getenv("MODEL"),
		})
	if err != nil {
		panic(err) // 检查模型创建错误
	}
	input := []*schema.Message{
		schema.SystemMessage("You are a helpful assistant."),
		schema.UserMessage("hello, who are you"),
	}
	// CompleteGenerate(ctx, model, input)
	StreamGenerate(ctx, model, input)
}
