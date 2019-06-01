package logger

import (
	"fmt"
	"time"

	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/tommytan/garen/configs"
)

// AliyunLoggerTest 阿里云日志服务
func AliyunLoggerTest() {
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = configs.GetConfiguration().LogEndpoint
	producerConfig.AccessKeyID = configs.GetConfiguration().LogAccessKey
	producerConfig.AccessKeySecret = configs.GetConfiguration().LogAccessKeySecret
	producerInstance := producer.InitProducer(producerConfig)
	//ch := make(chan os.Signal)
	//signal.Notify(ch)
	producerInstance.Start()
	for i := 0; i < 5; i++ {
		log := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"content": "test", "content2": fmt.Sprintf("%v %v", time.Now().Unix(), i)})
		err := producerInstance.SendLog("garen-go-test", "garen-go-test", "test-127.0.0.1", "topic", log)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("log sent to aliyun logger service")
	//if _, ok := <-ch; ok {
	//	fmt.Println("Get the shutdown signal and start to shut down")
	//	producerInstance.SafeClose()
	//}
	producerInstance.SafeClose()
}
