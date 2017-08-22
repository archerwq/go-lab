package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var filesNotToUpdate = []string{
	"/Users/qiangwang/dev/git/archerwq.github.io/_posts/2015-09-28-git-notes.md",
	"/Users/qiangwang/dev/git/archerwq.github.io/_posts/2015-10-09-java-logging-integrate-slf4j-and-logback.md",
	"/Users/qiangwang/dev/git/archerwq.github.io/_posts/2015-11-22-git-rebase-squash.md",
	"/Users/qiangwang/dev/git/archerwq.github.io/_posts/2016-01-04-docker-notes.md",
}

var titleMap = map[string]string{
	"wuzuolou_shuiku.jpg":              "2019/12/31/As5IP.jpg",
	"wuzuolou_shanlu.jpg":              "2019/12/31/Asp1h.jpg",
	"stored_program.jpg":               "2019/12/31/As9BI.jpg",
	"sonata_car.jpg":                   "2019/12/31/AsuR6.jpg",
	"river_inn_bigsur.jpg":             "2019/12/31/AsDgq.jpg",
	"nasa_visitor_center.jpg":          "2019/12/31/AsWEH.jpg",
	"moon_walk.jpg":                    "2019/12/31/AsJhd.jpg",
	"miyun_shuiku.jpg":                 "2019/12/31/AsCIf.jpg",
	"miyun_shuiku_bian.jpg":            "2019/12/31/Asw1o.jpg",
	"man_month.jpg":                    "2019/12/31/AsfA3.jpg",
	"light_house_beach.jpg":            "2019/12/31/AsNBK.jpg",
	"jekyll_and_docker.jpg":            "2019/12/31/Asqfj.jpg",
	"hurricane.jpg":                    "2019/12/31/AsA3U.jpg",
	"highway_route_one.jpg":            "2019/12/31/As8c0.jpg",
	"half_moon_bay.jpg":                "2019/12/31/As6Dv.jpg",
	"googleplex.jpg":                   "2019/12/31/AivhG.jpg",
	"google_1.jpg":                     "2019/12/31/Ail4T.jpg",
	"google_0.jpg":                     "2019/12/31/Aih1A.jpg",
	"git_rebase_2.jpg":                 "2019/12/31/Ai22s.jpg",
	"git_rebase_1.jpg":                 "2019/12/31/AiscC.jpg",
	"fuhuali_bar.jpg":                  "2019/12/31/AiTeu.jpg",
	"differential_analyzer.jpg":        "2019/12/31/AiZSX.jpg",
	"car_transfer.jpg":                 "2019/12/31/AiPDq.jpg",
	"car_transfer_process.jpg":         "2019/12/31/AigLH.jpg",
	"bixby.jpg":                        "2019/12/31/Aic0d.jpg",
	"bixby_2.jpg":                      "2019/12/31/AiSGf.jpg",
	"aifeibao_xiaozhen.jpg":            "2019/12/31/Ai78o.jpg",
	"aifeibao_putaoyuan.jpg":           "2019/12/31/AikX3.jpg",
	"ai_robot.jpg":                     "2019/12/31/AiRHK.jpg",
	"zhuhaihangzhan.jpg":               "2019/12/31/AiQYg.jpg",
	"zh_sunny_cloud.jpg":               "2019/12/31/Aiysj.jpg",
	"zh_sunny_cloud_2.jpg":             "2019/12/31/AinSU.jpg",
	"zh_nightfall.jpg":                 "2019/12/31/Ai5W0.jpg",
	"zh_nightfall_4.jpg":               "2019/12/31/Ai1jv.jpg",
	"zh_nightfall_3.jpg":               "2019/12/31/AiG0G.jpg",
	"zh_nightfall_2.jpg":               "2019/12/31/Ai99T.jpg",
	"zh_nightfall_1.jpg":               "2019/12/31/Aiu8A.jpg",
	"view_from_zh_home.jpg":            "2019/12/31/AiUXn.jpg",
	"zh_favorite_flower.jpg":           "2019/12/31/AiEyB.jpg",
	"zh_cloudy.jpg":                    "2019/12/31/AiDYN.jpg",
	"reading.jpg":                      "2019/12/31/AiKsR.jpg",
	"redwood.jpg":                      "2019/12/31/AiJ7s.jpg",
	"qiao_island.jpg":                  "2019/12/31/AiCWC.jpg",
	"qiao_bridge_sunset.jpg":           "2019/12/31/AifjS.jpg",
	"green_park.jpg":                   "2019/12/31/AimVh.jpg",
	"baishi_street_1.jpg":              "2019/12/31/AiAyu.jpg",
	"hongshulin.jpg":                   "2019/12/31/Ai8NI.jpg",
	"drink.jpg":                        "2019/12/31/Adxs6.jpg",
	"danjiang.jpg":                     "2019/12/31/Adv7p.jpg",
	"xjcy.jpg":                         "2019/12/31/AdlKX.jpg",
	"the_ballad_of_buster_scruggs.jpg": "2019/12/31/AdjZH.jpg",
	"stumbling_on_happiness.jpg":       "2019/12/31/Adzud.jpg",
	"night_cycling.jpg":                "2019/12/31/Adt6f.jpg",
	"headlights_in_the_distance.jpg":   "2019/12/31/Ad3Vo.jpg",
	"yelidao.jpg":                      "2019/12/30/Ad2n3.jpg",
	"us_trip_20171028.png":             "2019/12/31/As1mu.png",
	"slf4j_integration.png":            "2019/12/31/Asawp.png",
	"sanmateo_bigsur.png":              "2019/12/31/AsE3X.png",
	"knuth_words.png":                  "2019/12/31/AsFQg.png",
	"git_workspace.png":                "2019/12/31/AiLAn.png",
	"git_repository_objects.png":       "2019/12/31/AizeB.png",
	"git_rebase_before.png":            "2019/12/31/AitQN.png",
	"git_rebase_after.png":             "2019/12/31/AirfR.png",
	"git_merge_before.png":             "2019/12/31/AiiDS.png",
	"git_merge_after.png":              "2019/12/31/AibLt.png",
	"git_commits.png":                  "2019/12/31/AiB4D.png",
	"git_branch.png":                   "2019/12/31/AieGP.png",
	"git_branch_strategy.png":          "2019/12/31/AiX8h.png",
	"docker.png":                       "2019/12/31/AiOHI.png",
	"docker_vs_vm.png":                 "2019/12/31/AiIf6.png",
	"docker_image_layers.png":          "2019/12/31/Ai02p.png",
	"qiao_beach.png":                   "2019/12/31/AiYZt.png",
	"thinking_in_bets.png":             "2019/12/31/AdLjq.png",
}

var postsFolder = "/Users/qiangwang/dev/git/archerwq.github.io/_posts"

func main() {
	var files []string
	err := filepath.Walk(postsFolder, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files[1:] {
		toUpdate := true
		for _, f := range filesNotToUpdate {
			if file == f {
				toUpdate = false
				break
			}
		}
		if toUpdate {
			fmt.Println(file)
			updateBlog(file)
		}
	}
}

func updateBlog(path string) {
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)

	newText := text
	for k, v := range titleMap {
		newText = strings.Replace(newText, k, v, -1)
	}

	if err := ioutil.WriteFile(path, []byte(newText), 0644); err != nil {
		log.Fatal(err)
	}
}
