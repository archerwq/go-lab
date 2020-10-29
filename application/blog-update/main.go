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
	// "/Users/qiangwang/dev/git/archerwq.github.io/_posts/2015-09-28-git-notes.md",
	// "/Users/qiangwang/dev/git/archerwq.github.io/_posts/2015-10-09-java-logging-integrate-slf4j-and-logback.md",
	// "/Users/qiangwang/dev/git/archerwq.github.io/_posts/2015-11-22-git-rebase-squash.md",
	// "/Users/qiangwang/dev/git/archerwq.github.io/_posts/2016-01-04-docker-notes.md",
}

var titleMap = map[string]string{
	"2019/12/31/As8c0.jpg": "highway_route_one.jpg",
	"2019/12/31/AivhG.jpg": "googleplex.jpg",
	"2019/12/31/Aiu8A.jpg": "zh_nightfall_1.jpg",
	"2019/12/31/AiLAn.png": "git_workspace.png",
	"2019/12/31/Asw1o.jpg": "miyun_shuiku_bian.jpg",
	"2019/12/31/AiSGf.jpg": "bixby_2.jpg",
	"2019/12/31/AiKsR.jpg": "reading.jpg",
	"2019/12/31/Ai02p.png": "docker_image_layers.png",
	"2019/12/31/AsWEH.jpg": "nasa_visitor_center.jpg",
	"2019/12/31/As6Dv.jpg": "half_moon_bay.jpg",
	"2019/12/31/Aiysj.jpg": "zh_sunny_cloud.jpg",
	"2019/12/31/AiDYN.jpg": "zh_cloudy.jpg",
	"2019/12/31/AsE3X.png": "sanmateo_bigsur.png",
	"2019/12/31/AitQN.png": "git_rebase_before.png",
	"2019/12/31/Ail4T.jpg": "google_1.jpg",
	"2019/12/31/Ai78o.jpg": "aifeibao_xiaozhen.jpg",
	"2019/12/31/AimVh.jpg": "green_park.jpg",
	"2019/12/31/AiIf6.png": "docker_vs_vm.png",
	"2019/12/31/Asqfj.jpg": "jekyll_and_docker.jpg",
	"2019/12/31/AiCWC.jpg": "qiao_island.jpg",
	"2019/12/31/Adxs6.jpg": "drink.jpg",
	"2019/12/31/AdjZH.jpg": "the_ballad_of_buster_scruggs.jpg",
	"2019/12/31/Adzud.jpg": "stumbling_on_happiness.jpg",
	"2019/12/31/AizeB.png": "git_repository_objects.png",
	"2019/12/31/Aih1A.jpg": "google_0.jpg",
	"2019/12/31/AiscC.jpg": "git_rebase_1.jpg",
	"2019/12/31/AigLH.jpg": "car_transfer_process.jpg",
	"2019/12/31/AiEyB.jpg": "zh_favorite_flower.jpg",
	"2019/12/31/AsNBK.jpg": "light_house_beach.jpg",
	"2019/12/31/Ai5W0.jpg": "zh_nightfall.jpg",
	"2019/12/31/AdlKX.jpg": "xjcy.jpg",
	"2019/12/30/Ad2n3.jpg": "yelidao.jpg",
	"2019/12/31/Asawp.png": "slf4j_integration.png",
	"2019/12/31/Asp1h.jpg": "wuzuolou_shanlu.jpg",
	"2019/12/31/AsJhd.jpg": "moon_walk.jpg",
	"2019/12/31/AsCIf.jpg": "miyun_shuiku.jpg",
	"2019/12/31/AsfA3.jpg": "man_month.jpg",
	"2019/12/31/AinSU.jpg": "zh_sunny_cloud_2.jpg",
	"2019/12/31/AiX8h.png": "git_branch_strategy.png",
	"2019/12/31/As9BI.jpg": "stored_program.jpg",
	"2019/12/31/AiPDq.jpg": "car_transfer.jpg",
	"2019/12/31/AikX3.jpg": "aifeibao_putaoyuan.jpg",
	"2019/12/31/AifjS.jpg": "qiao_bridge_sunset.jpg",
	"2019/12/31/AiOHI.png": "docker.png",
	"2019/12/31/AsDgq.jpg": "river_inn_bigsur.jpg",
	"2019/12/31/AsA3U.jpg": "hurricane.jpg",
	"2019/12/31/AiUXn.jpg": "view_from_zh_home.jpg",
	"2019/12/31/Ai8NI.jpg": "hongshulin.jpg",
	"2019/12/31/AiRHK.jpg": "ai_robot.jpg",
	"2019/12/31/AiG0G.jpg": "zh_nightfall_3.jpg",
	"2019/12/31/AiAyu.jpg": "baishi_street_1.jpg",
	"2019/12/31/AsFQg.png": "knuth_words.png",
	"2019/12/31/AdLjq.png": "thinking_in_bets.png",
	"2019/12/31/Ai22s.jpg": "git_rebase_2.jpg",
	"2019/12/31/Aic0d.jpg": "bixby.jpg",
	"2019/12/31/AiJ7s.jpg": "redwood.jpg",
	"2019/12/31/AiB4D.png": "git_commits.png",
	"2019/12/31/AsuR6.jpg": "sonata_car.jpg",
	"2019/12/31/AiTeu.jpg": "fuhuali_bar.jpg",
	"2019/12/31/AiZSX.jpg": "differential_analyzer.jpg",
	"2019/12/31/AiQYg.jpg": "zhuhaihangzhan.jpg",
	"2019/12/31/As1mu.png": "us_trip_20171028.png",
	"2019/12/31/Ad3Vo.jpg": "headlights_in_the_distance.jpg",
	"2019/12/31/Adv7p.jpg": "danjiang.jpg",
	"2019/12/31/Adt6f.jpg": "night_cycling.jpg",
	"2019/12/31/AibLt.png": "git_merge_after.png",
	"2019/12/31/AiYZt.png": "qiao_beach.png",
	"2019/12/31/As5IP.jpg": "wuzuolou_shuiku.jpg",
	"2019/12/31/Ai1jv.jpg": "zh_nightfall_4.jpg",
	"2019/12/31/Ai99T.jpg": "zh_nightfall_2.jpg",
	"2019/12/31/AirfR.png": "git_rebase_after.png",
	"2019/12/31/AiiDS.png": "git_merge_before.png",
	"2019/12/31/AieGP.png": "git_branch.png",
}

var newMap = map[string]string{
	"stumbling_on_happiness.jpeg":      "5f959d911cd1bbb86ba55af2.jpg",
	"sanmateo_bigsur.png":              "5f959f851cd1bbb86ba5d807.png",
	"docker_vs_vm.png":                 "5f959f1b1cd1bbb86ba5baa6.png",
	"hurricane.jpg":                    "5f959f4e1cd1bbb86ba5ca1d.jpg",
	"git_workspace.png":                "5f959f1c1cd1bbb86ba5bae9.png",
	"git_rebase_1.jpeg":                "5f959f1b1cd1bbb86ba5bad0.jpg",
	"kso_breakfast.jpg":                "5f959e0d1cd1bbb86ba57687.jpg",
	"light_house_beach.jpg":            "5f959f4e1cd1bbb86ba5ca37.jpg",
	"highway_route_one.jpg":            "5f959f4e1cd1bbb86ba5ca17.jpg",
	"google_1.jpg":                     "5f95952b1cd1bbb86ba3529b.jpg",
	"git_commits.png":                  "5f959f1b1cd1bbb86ba5babe.png",
	"differential_analyzer.jpg":        "5f9593ad1cd1bbb86ba2fa7e.jpg",
	"zhuhaihangzhan.jpeg":              "5f959ef71cd1bbb86ba5b0c6.jpg",
	"aifeibao_putaoyuan.jpg":           "5f959ef71cd1bbb86ba5b0d2.jpg",
	"view_from_zh_home.jpg":            "5f9595f41cd1bbb86ba3837e.jpg",
	"slf4j_integration.png":            "5f959f851cd1bbb86ba5d80d.png",
	"knuth_words.jpg":                  "5f959f4e1cd1bbb86ba5ca2f.png",
	"the_ballad_of_buster_scruggs.jpg": "5f959d911cd1bbb86ba55af7.jpg",
	"jekyll_and_docker.jpg":            "5f959f4e1cd1bbb86ba5ca27.jpg",
	"drink.jpg":                        "5f959d911cd1bbb86ba55adc.jpg",
	"wuzuolou_shuiku.jpg":              "5f959b671cd1bbb86ba4d841.jpg",
	"stored_program.jpg":               "5f959f851cd1bbb86ba5d817.jpg",
	"git_rebase_after.png":             "5f959f1c1cd1bbb86ba5bad9.png",
	"car_transfer.jpg":                 "5f959f1b1cd1bbb86ba5ba93.jpg",
	"miyun_shuiku_bian.jpg":            "5f959f4f1cd1bbb86ba5ca50.jpg",
	"git_branch_strategy.png":          "5f959f1b1cd1bbb86ba5bab3.png",
	"yelidao.jpeg":                     "5f959d911cd1bbb86ba55b05.jpg",
	"miyun_shuiku.jpg":                 "5f959f4f1cd1bbb86ba5ca53.jpg",
	"qiao_island.jpg":                  "5f959e0d1cd1bbb86ba576a4.jpg",
	"qiao_beach.png":                   "5f959e0d1cd1bbb86ba57693.png",
	"redwood.jpg":                      "5f959e0d1cd1bbb86ba576ad.jpg",
	"headlights_in_the_distance.jpg":   "5f959d911cd1bbb86ba55ae2.jpg",
	"sonata_car.jpg":                   "5f959f851cd1bbb86ba5d813.jpg",
	"baishi_street_1.jpg":              "5f959e0d1cd1bbb86ba5767a.jpg",
	"nasa_visitor_center.jpg":          "5f959f851cd1bbb86ba5d7fc.jpg",
	"qiao_bridge_sunset.jpg":           "5f959e0d1cd1bbb86ba5769e.jpg",
	"bixby.jpg":                        "5f959ef71cd1bbb86ba5b10a.jpg",
	"kso_office.jpg":                   "5f959e0d1cd1bbb86ba5768e.jpg",
	"git_repository_objects.png":       "5f959f1c1cd1bbb86ba5bae5.png",
	"aifeibao_xiaozhen.jpg":            "5f959ef71cd1bbb86ba5b0d6.jpg",
	"night_cycling.jpg":                "5f959d911cd1bbb86ba55aed.jpg",
	"thinking_in_bets.png":             "5f959d911cd1bbb86ba55afb.png",
	"googleplex.jpg":                   "5f959f4e1cd1bbb86ba5ca0c.jpg",
	"fuhuali_bar.jpg":                  "5f959f1b1cd1bbb86ba5baae.jpg",
	"half_moon_bay.jpg":                "5f959f4e1cd1bbb86ba5ca11.jpg",
	"git_merge_before.png":             "5f959f1b1cd1bbb86ba5bac9.png",
	"ai_robot.jpg":                     "5f959ef71cd1bbb86ba5b0cb.jpg",
	"zh_nightfall_3.jpg":               "5f959ec41cd1bbb86ba5a515.jpg",
	"green_park.jpg":                   "5f959e0d1cd1bbb86ba57681.jpg",
	"danjiang.jpg":                     "5f959d911cd1bbb86ba55ad9.jpg",
	"wuzuolou_shanlu.jpg":              "5f959f851cd1bbb86ba5d81f.jpg",
	"river_inn_bigsur.jpg":             "5f959f851cd1bbb86ba5d803.jpg",
	"xjcy.jpg":                         "5f959d911cd1bbb86ba55afe.jpg",
	"zh_cloudy.jpg":                    "5f959ec41cd1bbb86ba5a4e2.jpg",
	"zh_nightfall_1.jpg":               "5f959ec41cd1bbb86ba5a502.jpg",
	"zh_nightfall_2.jpg":               "5f959ec41cd1bbb86ba5a50e.jpg",
	"git_rebase_2.jpeg":                "5f959f1b1cd1bbb86ba5bad3.jpg",
	"zh_nightfall.jpg":                 "5f959ef71cd1bbb86ba5b0b6.jpg",
	"git_rebase_before.png":            "5f959f1c1cd1bbb86ba5bae0.png",
	"docker_image_layers.png":          "5f959f1b1cd1bbb86ba5baa0.png",
	"docker.png":                       "5f959f1b1cd1bbb86ba5baab.png",
	"git_branch.png":                   "5f959f1b1cd1bbb86ba5bab8.png",
	"zh_sunny_cloud.jpg":               "5f959ef71cd1bbb86ba5b0c0.jpg",
	"bixby_2.jpg":                      "5f959ef71cd1bbb86ba5b0ff.jpg",
	"car_transfer_process.jpg":         "5f959ef71cd1bbb86ba5b10e.jpg",
	"hongshulin.jpg":                   "5f959d911cd1bbb86ba55ae6.jpg",
	"us_trip_20171028.png":             "5f959f851cd1bbb86ba5d81b.png",
	"git_merge_after.png":              "5f959f1b1cd1bbb86ba5bac5.png",
	"reading.jpeg":                     "5f959e0d1cd1bbb86ba576a8.jpg",
	"moon_walk.jpg":                    "5f959f851cd1bbb86ba5d7f5.jpg",
	"google_0.jpg":                     "5f959f4e1cd1bbb86ba5ca02.jpg",
	"zh_favorite_flower.jpg":           "5f959ec41cd1bbb86ba5a4fe.jpg",
	"zh_nightfall_4.jpg":               "5f959ec41cd1bbb86ba5a51d.jpg",
	"man_month.jpg":                    "5f959f4e1cd1bbb86ba5ca3c.jpg",
	"zh_sunny_cloud_2.jpg":             "5f959ef71cd1bbb86ba5b0ba.jpg",
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
		if newURL, ok := newMap[v]; ok {
			newText = strings.Replace(newText, k, newURL, -1)
		}
	}

	if err := ioutil.WriteFile(path, []byte(newText), 0644); err != nil {
		log.Fatal(err)
	}
}
