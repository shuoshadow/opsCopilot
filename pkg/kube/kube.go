package kube

import (
	"context"
	"fmt"
	"time"

	opsv1 "github.com/shaowenchen/ops/api/v1"
	"github.com/shaowenchen/ops/pkg/constants"
	"github.com/shaowenchen/ops/pkg/log"
	"github.com/shaowenchen/ops/pkg/utils"
	v1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

func Script(logger *log.Logger, client *kubernetes.Clientset, node v1.Node, option ScriptOption) (err error) {
	namespacedName, err := utils.GetOrCreateNamespacedName(client, constants.OpsNamespace, fmt.Sprintf("script-%s", time.Now().Format("2006-01-02-15-04-05")))
	if err != nil {
		logger.Error.Println(err)
	}
	pod, err := RunScriptOnNode(client, node, namespacedName, option.Image, option.Content)
	if err != nil {
		logger.Error.Println(err)
	}
	GetPodLog(context.TODO(), logger, client, pod)
	return
}

func File(logger *log.Logger, client *kubernetes.Clientset, node v1.Node, option FileOption) (err error) {
	namespacedName, err := utils.GetOrCreateNamespacedName(client, constants.OpsNamespace, fmt.Sprintf("file-%s", time.Now().Format("2006-01-02-15-04-05")))
	if err != nil {
		logger.Error.Println(err)
	}
	pod, err := DownloadFileOnNode(client, node, namespacedName, option.Image, option.RemoteFile, option.LocalFile)
	if err != nil {
		logger.Error.Println(err)
	}
	GetPodLog(context.TODO(), logger, client, pod)
	return
}

func GetPodLog(ctx context.Context, logger *log.Logger, client *kubernetes.Clientset, pod *v1.Pod) (log string, err error) {
	for range time.Tick(time.Second * 3) {
		select {
		default:
			pod, err = client.CoreV1().Pods(pod.Namespace).Get(ctx, pod.Name, metav1.GetOptions{})
			if utils.IsPendingPod(pod) {
				continue
			}
			log, err = utils.GetPodLog(ctx, client, pod.Namespace, pod.Name)
			if err != nil {
				logger.Error.Println(err)
				return
			}
			logger.Info.Println(log)
			if utils.IsStopedPod(pod) {
				return
			}
		}
	}
	return
}

func GetNodes(logger *log.Logger, client *kubernetes.Clientset, option KubeOption) (nodeList []v1.Node, err error) {
	nodes, err := utils.GetAllNodesByClient(client)
	if err != nil {
		logger.Error.Println(err)
		return
	}
	if len(option.NodeName) > 0 {
		for _, node := range nodes.Items {
			if node.Name == option.NodeName {
				nodeList = append(nodeList, node)
			}
		}
	}
	if option.All {
		nodeList = nodes.Items
	}
	return
}

func GetOpsClient(logger *log.Logger, restConfig *rest.Config) (client runtimeClient.Client, err error) {
	scheme, err := opsv1.SchemeBuilder.Build()
	if err != nil {
		return
	}

	client, err = runtimeClient.New(restConfig, runtimeClient.Options{Scheme: scheme})
	if err != nil {
		return
	}
	return
}