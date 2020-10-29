package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
)

const (
	// srcURLPref  = "https://cdn.img.wenhairu.com/images/"
	srcURLPref  = "https://img1.wenhairu.com/images/"
	newURLPref  = "https://pic.imgdb.cn/item/"
	uploadAPI   = "https://api.superbed.cn/upload"
	postsFolder = "/Users/qiangwang/dev/git/archerwq.github.io/_posts"
)

var srcImgPaths = []string{
	// "2020/07/12/fBD03.png",
	// "2019/12/31/As5IP.jpg",
	// "2019/12/31/Asp1h.jpg",
	// "2019/12/31/As9BI.jpg",
	// "2019/12/31/AsuR6.jpg",
	// "2019/12/31/AsDgq.jpg",
	// "2019/12/31/AsWEH.jpg",
	// "2019/12/31/AsJhd.jpg",
	// "2019/12/31/AsCIf.jpg",
	// "2019/12/31/Asw1o.jpg",
	// "2019/12/31/AsfA3.jpg",
	// "2019/12/31/AsNBK.jpg",
	// "2019/12/31/Asqfj.jpg",
	// "2019/12/31/AsA3U.jpg",
	// "2019/12/31/As8c0.jpg",
	// "2019/12/31/As6Dv.jpg",
	// "2019/12/31/AivhG.jpg",
	// "2019/12/31/Ail4T.jpg",
	// "2019/12/31/Aih1A.jpg",
	// "2019/12/31/Ai22s.jpg",
	// "2019/12/31/AiscC.jpg",
	// "2019/12/31/AiTeu.jpg",
	// "2019/12/31/AiZSX.jpg",
	// "2019/12/31/AiPDq.jpg",
	// "2019/12/31/AigLH.jpg",
	// "2019/12/31/Aic0d.jpg",
	// "2019/12/31/AiSGf.jpg",
	// "2019/12/31/Ai78o.jpg",
	// "2019/12/31/AikX3.jpg",
	// "2019/12/31/AiRHK.jpg",
	// "2019/12/31/AiQYg.jpg",
	// "2019/12/31/Aiysj.jpg",
	// "2019/12/31/AinSU.jpg",
	// "2019/12/31/Ai5W0.jpg",
	// "2019/12/31/Ai1jv.jpg",
	// "2019/12/31/AiG0G.jpg",
	// "2019/12/31/Ai99T.jpg",
	// "2019/12/31/Aiu8A.jpg",
	"2019/12/31/AiUXn.jpg",
	// "2019/12/31/AiEyB.jpg",
	// "2019/12/31/AiDYN.jpg",
	// "2019/12/31/AiKsR.jpg",
	// "2019/12/31/AiJ7s.jpg",
	// "2019/12/31/AiCWC.jpg",
	// "2019/12/31/AifjS.jpg",
	// "2019/12/31/AimVh.jpg",
	// "2019/12/31/AiAyu.jpg",
	// "2019/12/31/Ai8NI.jpg",
	// "2019/12/31/Adxs6.jpg",
	// "2019/12/31/Adv7p.jpg",
	// "2019/12/31/AdlKX.jpg",
	// "2019/12/31/AdjZH.jpg",
	// "2019/12/31/Adzud.jpg",
	// "2019/12/31/Adt6f.jpg",
	// "2019/12/31/Ad3Vo.jpg",
	// "2019/12/30/Ad2n3.jpg",
	// "2019/12/31/As1mu.png",
	// "2019/12/31/Asawp.png",
	// "2019/12/31/AsE3X.png",
	// "2019/12/31/AsFQg.png",
	// "2019/12/31/AiLAn.png",
	// "2019/12/31/AizeB.png",
	// "2019/12/31/AitQN.png",
	// "2019/12/31/AirfR.png",
	// "2019/12/31/AiiDS.png",
	// "2019/12/31/AibLt.png",
	// "2019/12/31/AiB4D.png",
	// "2019/12/31/AieGP.png",
	// "2019/12/31/AiX8h.png",
	// "2019/12/31/AiOHI.png",
	// "2019/12/31/AiIf6.png",
	// "2019/12/31/Ai02p.png",
	// "2019/12/31/AiYZt.png",
	// "2019/12/31/AdLjq.png",
	// "2020/07/12/fBD03.png",
}

type uploadResult struct {
	Err    int               `json:"err"`
	URL    string            `json:"url"`
	URLMap map[string]string `json:"urls"`
	Msg    string            `json:"msg"`
}

var client *http.Client
var token = flag.String("token", "", "token for upload")

func init() {
	client = &http.Client{}
}

func main() {
	flag.Parse()

	pathMap := make(map[string]string, len(srcImgPaths))
	for _, path := range srcImgPaths {
		log.Printf("===> %s\n", path)
		srcURL := srcURLPref + path
		result, err := uplaodWithShell(srcURL)
		if err != nil {
			panic(fmt.Errorf("failed to upload %s: %v", srcURL, err))
		}
		destPath := strings.Replace(result.URL, newURLPref, "", -1)
		pathMap[path] = destPath
		log.Printf("%s -> %s\n", path, destPath)
	}
}

func upload(imgSrcURL string) (*uploadResult, error) {
	log.Printf("uploading %s\n", imgSrcURL)

	data := url.Values{}
	data.Set("src", imgSrcURL)

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s?token=%s", uploadAPI, *token),
		strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do the request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	log.Println(string(body))

	var result uploadResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %v", err)
	}
	return &result, nil
}

func uplaodWithShell(imgSrcURL string) (*uploadResult, error) {
	log.Printf("uploading %s\n", imgSrcURL)

	app := "curl"
	data := fmt.Sprintf("'src=%s'", imgSrcURL)
	target := fmt.Sprintf("%s\\?token\\=%s", uploadAPI, *token)

	cmd := exec.Command(app, "-d", data, target)
	fmt.Println(cmd)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to curl: %v", err)
	}
	fmt.Println(string(output))

	var result uploadResult
	if err := json.Unmarshal(output, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %v", err)
	}
	return &result, nil
}
