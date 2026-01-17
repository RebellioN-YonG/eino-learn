package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
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

func TemplateMsg(ctx context.Context, input map[string]any) []*schema.Message {
	// create template
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}"),
		&schema.Message{
			Role:    schema.User,
			Content: "我入门了， {task}",
		},
	)

	msg, err := template.Format(ctx, input)

	if err != nil {
		panic(err)
	}

	return msg
}
func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	model, err := ark.NewChatModel(ctx,
		&ark.ChatModelConfig{
			APIKey: os.Getenv("ARK_API_KEY"),
			Model:  os.Getenv("MODEL"),
		})
	// create template
	// template := prompt.FromMessages(schema.FString,
	// 	schema.SystemMessage("你是一个{role}"),
	// 	&schema.Message{
	// 		Role:    schema.User,
	// 		Content: "我入门了， {task}",
	// 	},
	// )
	// prepare params
	params := map[string]any{
		"role": "福利鸡",
		"task": "和我做",
	}
	msg := TemplateMsg(ctx, params)
	// msg, err := template.Format(ctx, params)

	// if err != nil {
	// panic(err)
	// }
	CompleteGenerate(ctx, model, msg)
}
