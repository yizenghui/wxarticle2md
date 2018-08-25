package wxarticle2md

import (
	"testing"
)

var (
	htmlx = []string{
		`<div>
<p><img data-ratio="0.09477521263669501" data-type="png" data-s="300,640" data-w="1646" src="https://mmbiz.qpic.cn/mmbiz_png/8xT0JYxK1STkibtsjBSq5oX6EuOhG1na3baJwpAOBWnLJUEVkpAegOibaKoibMibVuskQrnWXnDm7EiaUvjsv0iacdGw/0?wx_fmt=png"/></p><p><br/></p><p><span>我在马路边，捡到一分钱。怎么办？</span></p><p><span>一个48秒的小魔术▼▼▼</span></p><p><br/></p><p><iframe data-vidtype="2" frameborder="0" allowfullscreen="" data-ratio="1.7647058823529411" data-w="480" data-src="https://v.qq.com/iframe/preview.html?vid=j0555ej6gau&amp;width=500&amp;height=375&amp;auto=0"></iframe></p><p><br/></p><p><span>开会、挤地铁、接孩子放学，月底流量不足又找不到wifi的话，</span><span>就先看精简图片版吧。</span><br/></p><p><br/></p><p><img data-ratio="0.6208333333333333" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5P19MPViau5ADAuX7q5kMX0BWarTRDD8FicPc7gVReVNfiaYj0MVTmgmdw/0?wx_fmt=jpeg"/></p><p><img data-ratio="0.5777777777777777" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5uxmbTiaQOXHOjwWvHzfMYv8SaElFbSRkXrflRz0d9D6iarh0ibYymey8w/0?wx_fmt=jpeg"/></p><p><img data-ratio="0.575" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5Elv4G2fzAtN33dmvgNJxlzAxQvj5qWs2gdDHYX9z8ouu8vIfZEr79Q/0?wx_fmt=jpeg"/></p><p><img data-ratio="0.5875" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5gxicnCGib8ZTDgiciaUicHF6djzNd7AYc48Zdcv1nRcRH4tVaUcUhPvaReA/0?wx_fmt=jpeg"/></p><p><img data-ratio="0.5819444444444445" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5YQcoElk2ib8jlOgNSG08hnJLJYorcoJEh2Zw1MyEiaFuDAiajS7Wcyafw/0?wx_fmt=jpeg"/></p><p><img data-ratio="0.5847222222222222" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5IgoYRXqibaNSbNv42mpoeJS535NsDUTdQeyRlKHIibB3Npicyd7wb4h3A/0?wx_fmt=jpeg"/></p><p><img data-ratio="0.5638888888888889" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5gQHibXwibwEcDTBaYzuZn87nric1JG9kyia7w0ohiaINJM3B2ohTdzS8ojQ/0?wx_fmt=jpeg"/></p><p><img data-ratio="0.5805555555555556" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5fwabpvSZpBLdZhqQsSY0xiaOsJeUL3DgkvQrPceonj8Z5ARZKricPJww/0?wx_fmt=jpeg"/></p><p><img data-ratio="0.5847222222222222" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe50pz3jHdIB6lvu3TAibLCJ4xEbwibUSA41rFxTBMtsibwe6HU3pUXkjwxA/0?wx_fmt=jpeg"/></p><p><br/></p><p><span><strong><span>有奖互个动</span></strong></span></p><p><span>可以直接刷二维码进出站的是哪条线？</span><span>评论区回答有点厉害的几位粉丝，发狗。</span></p><p><span>*请注意！我台有奖互动截至时间为：每周五上午10点钟整，本台不支持刷赞哦。</span></p><p><br/></p><p><span>戳</span><span><strong><span>阅读原文</span></strong></span><span><strong><span></span></strong></span><span><br/>京东支付教你一分钱坐北京地铁。</span></p><p><span><br/></span></p><p><img data-ratio="0.21113243761996161" data-type="png" data-s="300,640" data-w="1563" src="https://mmbiz.qpic.cn/mmbiz_png/8xT0JYxK1SQpcwChfZT3Ip3knKQqHR8GkNOM8fqeSEwkz2LoIqbctjplIP96ESyeAhmMKNgJuVSTmKq3RM958Q/0?wx_fmt=png"/></p><p><br/></p>	
</div>
`, `<html><head></head><body>
		<div>
							<p><br/></p><section><section><section><section data-width="90%">
							<section><section><section><section><section><section><section><section>
							<section><section><section><span>听说关注这个号，会成为温暖的人</span></section>
							<section><span>
							<img data-w="700" data-ratio="0.5614285714285714" src="https://mmbiz.qpic.cn/mmbiz_gif/icB0yCLh6LJuxzHw2NibZHgHyzfLx8o45vall8J2vHFIuPR4iaSIm9MYNOX3VaVXkL4RbF4DO1cTDb3zpnFicFLjtw/640?wx_fmt=gif" data-type="gif"/></span></section><section>
							<section></section><section></section><span>◆</span><span> </span><span>◆</span><span> </span>
							<span>◆</span><br/></section><section><span><p><span></span></p>
							</span></section></section></section></section></section></section>
							</section></section></section></section><p><span>「“水深则流缓，语迟则人贵。”」</span></p>
							<p><span><span>     </span></span><span> from温暖-</span><span><span>第748篇原创</span></span><span>      </span></p><section><section><span><br/>
							</span></section><section><span>    每早六点   </span></section></section><section><section><span><br/>
							</span></section><section><span>遇见每一个     向阳而生的你</span></section>
							</section><p>
							
	<mpvoice frameborder="0" voice_encode_fileid="MzIzMzA3NDg0Nl8yNjQ5NTk4MTg0" src="/cgi-bin/readtemplate?t=tmpl/audio_tmpl&amp;name=%E6%89%A7%E7%9D%80&amp;play_length=05:30" isaac2="0" low_size="618.01" source_size="618" high_size="2573.85" name="执着" play_length="330000"></mpvoice>

							</p>
							
							</section></section></section>
							</section></section><section>
							<section powered-by="xiumi.us"><section><section></section></section></section></section><p><span><strong>◆</strong></span><strong><span data-txtless="spin" data-txtlessp="120"> </span><span>◆</span><span> </span><span>◆</span></strong></p><p><span>文 | 李思圆</span></p>
							<p><img data-w="400" data-ratio="0.15" src="https://mmbiz.qpic.cn/mmbiz/icB0yCLh6LJu91ia36ryqgAMfr1ahXFuxCc8gsAdvDSM0ibFNrmp11uaP5CDib6T1WRibrWOjW4YcPCo6uib2dVE3UVg/640?wx_fmt=png" data-type="png"/></p><p><br/></p><section><section powered-by="xiumi.us"><section><section>
							<p><img data-w="1280" data-copyright="0" data-ratio="0.13984375" data-s="300,640" src="https://mmbiz.qpic.cn/mmbiz_jpg/icB0yCLh6LJs2x3AQXx794PzvDYpLYlicJAaEZiaUoHHYG2vibesGneWW9QfSodibaFrib6o7AIaOauw5ia5DCDYx2U9A/640?wx_fmt=jpeg" data-type="jpeg"/></p><p><br/></p><p>朋友昨天给我打来电话，说她今天中午感觉吃了一顿鸿门宴，整个人惊恐万分。</p><p><br/></p><p>我连忙问她怎么了。</p><p><br/></p><p>她说，她们公司财务总监今天主动邀请她一起吃饭，而且在路上还主动跟她挽着手走。</p><p><br/></p><p>我还责怪朋友，得了好处还卖乖。朋友说你不知道她这个人平时的为人，她老是爱发脾气，这么谦和地待人还是很少见的，真不太能适应，然后她就给我讲了她们之间的事儿。</p><p><br/></p><p>朋友刚到公司那会儿，经常拿着各类票据要找财务总监报账，可只要稍有不对的地方，财务总监就像一座火山，瞬间爆发。前一秒也许她还笑脸盈盈，下一秒就如一头发怒的狮子，着实让人震惊。</p><p><br/></p><p>原本想着本来自己是职场新人而且自己犯的错，理应承受这些。于是就没放心上。</p><p><br/></p><p>可还有一次财务报销程序换了流程，她事前不知道，还是以老方法去报账，结果财务总监又发火了，她看了以后直接把报销单扔在地上，脸气得通红，直接甩朋友一句，报不了。</p><p><br/></p><p>朋友很无辜且小心翼翼地问道：“是哪儿错了？错了我马上改。”谁知财务总监根反正就板着脸，只顾着发脾气。</p><p><br/></p><p>职场上同事之间，最容易因为一些工作上的小问题发脾气，并且这些问题又不是什么大不了的事儿，可很多人就控制不住自己的情绪，或者只顾发脾气，思维只停留在错的地方，根本就不去想挽救和解决的办法。</p><p><br/></p><p>于是你怪我，我怪你，大家互相指责对方，最后事情没解决，双方脾气倒是越来越大。</p><p><br/></p><p>杜月笙曾说：头等人，有本事，没脾气；二等人，有本事，有脾气；末等人，没本事，大脾气。</p><p><br/></p><p><strong><span>生活中，一个无法控制自己情绪，且爱发脾气的人，其实就是修养不够。</span></strong></p><p><br/></p><p><img data-w="1280" data-copyright="0" data-ratio="0.13984375" data-s="300,640" src="https://mmbiz.qpic.cn/mmbiz_jpg/icB0yCLh6LJs2x3AQXx794PzvDYpLYlicJPiaUrFEkk6dGFjB1BhqbJJzojtZw5LibzyLd4odSSwydpFT390NubsGA/640?wx_fmt=jpeg" data-type="jpeg"/></p><p><br/></p><p>还有一个朋友是做行政工作的。最近她们公司因为业务发展需要，要更换公司名称，到了工商局去办理时，工作人员说需要今年银行出具的验资报告，或者律师事务所的审计报告。</p><p><br/></p><p>于是朋友回来拿了这些资料然后去办理，正办的时候，旁边有个人也跟她办理同样的业务。</p><p><br/></p><p>但不同的是，当窗口人员告诉她需要这些手续时，这个人瞬间就火冒三丈，她说：“我都问了我们公司，公司说这是很私密的资料，不允许对外出示。”</p><p><br/></p><p>朋友连忙帮着打圆场说到：“我们公司也是这样办的，也需要同样的手续。&#34;可那个人就一副，我不管，我今天就算资料不齐，也要办理的样子。</p><p><br/></p><p>无论她怎么生气，窗口人员依旧忍住怒火，耐心地一再向她解释，这是正常的工作流程，并不是有意为难她。</p><p><br/></p><p>直到朋友离开的时候，那个人就跟窗口人员吵了起来，她又是拍桌子，又是破口大骂，然后又是威胁到，我要找你们领导投诉你。</p><p><br/></p><p><img data-w="1080" data-copyright="0" data-ratio="0.5675925925925925" data-s="300,640" src="https://mmbiz.qpic.cn/mmbiz_jpg/icB0yCLh6LJvfBOuJv6s4wLRoD3SeFDXPY2Gug84vsjFPBuY3BoeljsdXQctL6kL9lBSfLGJqR3fAHGSboicKUIg/640?wx_fmt=jpeg" data-type="jpeg"/></p><p><br/></p><p>很多时候一些基层的窗口人员和办事人员之间发生的矛盾，很大程度上都归于双方任意一方脾气不好，态度不好，不好好说话。</p><p><br/></p><p>有部分蛮横不讲理的人，他们总以为向别人发了脾气就能解决问题，甚至是威胁恐吓，可脾气大的人，本事却不大，他们总以为自己是对的，总想发脾气把对方震慑住。</p><p><br/></p><p>脾气爆裂的人，经常因别人不合他意或一两句不遂心愿的话，便怒火中烧，当脾气冲昏了头脑，他们甚至不辩明是非，就一味地发脾气。人愤怒的一瞬间，智商就为零。</p><p><br/></p><p>但一个人的修养关键在于控制自己的情绪，用嘴伤人，用气势伤人，既不能解决实际问题，也是最愚蠢的一种行为。</p><p><br/></p><p><img data-w="1280" data-copyright="0" data-ratio="0.13984375" data-s="300,640" src="https://mmbiz.qpic.cn/mmbiz_jpg/icB0yCLh6LJs2x3AQXx794PzvDYpLYlicJk4gFSMBb5PCOBia7W81r2uFa7HGhib4W3doDXwTKsvlsgdFDSoiaOjHYA/640?wx_fmt=jpeg" data-type="jpeg"/></p><p><br/></p><p>有一次在餐馆吃饭时，无意间听到我隔壁的一个家长爸爸跟一群同行的人聊天。</p><p><br/></p><p>他说，女儿这学期期末考试成绩不理想，于是回家以后，他扇了女儿几耳光，还拳打脚踢。打了以后，又觉得有些不忍心，于是今天决定带着女儿去欢乐谷玩玩儿，弥补和修复一下父女之间的感情。</p><p><br/></p><p>可我仔细观察了那个正在吃饭的女儿，她全程吃饭都低着头，爸爸一说话，她都生害怕是自己犯了错，又要被罚，整个人的状态就像一只怯怯微微的小老鼠，随时都怕被身边脾气像老虎的爸爸给吃了。</p><p><br/></p><p>这时候就有人对他说，你看你脾气这么差，你女儿怕你啊。</p><p><br/></p><p>他说，其实当时就是控制不住情绪，打了她就后悔了。</p><p><br/></p><p>生活中，有多少这样的父母，他们自己都控制不好自己的情绪，孩子稍微犯错，就被父母大肆责罚，他们一味地发脾气和打骂，殊不知这样根本也没有达到预期的教育意义。</p><p><br/></p><p>凡事都爱发脾气的父母，他们总是教育孩子要有教养、有涵养、有修养，其实他们自己都做不到，就连他们自己都像一枚不定时炸弹，又怎能要求孩子遇事时处乱不惊，沉着镇定。</p><p><br/></p><p><img data-w="1080" data-copyright="0" data-ratio="0.5675925925925925" data-s="300,640" src="https://mmbiz.qpic.cn/mmbiz_jpg/icB0yCLh6LJvfBOuJv6s4wLRoD3SeFDXP2ABqLpqZsNCQiccv5kfG55wF9Gc8YbMWRIRxKDbfVW9gFrv7RWJBatg/640?wx_fmt=jpeg" data-type="jpeg"/></p><p><br/></p><p>好脾气其实是人生的一笔财富，而父母给予孩子最好的财富不是遗留多少的票子和房子，而是以身作则，自己拥有了一个好脾气，也让孩子耳濡目染，真正收获到这笔财富。</p><p><br/></p><p><img data-w="1280" data-copyright="0" data-ratio="0.13984375" data-s="300,640" src="https://mmbiz.qpic.cn/mmbiz_jpg/icB0yCLh6LJs2x3AQXx794PzvDYpLYlicJakr4dJdiaZa7bXSCo7JOdRIRp8qTTO0VFkVibvCktlxrTEmTWGoHVZ2A/640?wx_fmt=jpeg" data-type="jpeg"/></p><p><br/></p><p>无论是在生活、职场，还是家庭中，我们都要学会处理好与同事朋友、亲人伴侣的关系。而一段关系的好坏，常常取决于是否有一个有修养的人，用一副好脾气温和待人。</p><p><br/></p><p>很多时候，你会发现即便你再有理，一旦发了脾气赢了一时气焰又如何？</p><p><br/></p><p>别人在心里照样不服你，人最难的不是证明自己的对错，最难的是收获人心，让别人打心底承认你是个明智且有修养的人。</p><p><br/></p><p>明智的人很多，而有修养的人却很少。</p><p><br/></p><p>但凡在一个有修养的人，都具有一个相同的品格，那就是拥有一个好脾气。一个能控制住情绪且不乱发脾气的人，比拿下一座城池的人更强大。</p><p><br/></p><p><strong><span>水深则流缓，语迟则人贵。脾气好才能拥有好的修养，也更能收获好人缘。</span></strong></p><p><br/></p><p>常常有人问，修养是什么？</p><p><br/></p><p>修养不是学历高、门第高，也不是姿态高，而是：与人说话时细雨和风、和颜悦色；待人时，温文儒雅、宽仁厚德；处事时，和蔼可亲，讲道理也讲情理。</p><p><br/></p>
							<p>一个有修养的人，他懂得控制自己的情绪，在自己对的时候，给别人退步的空间。</p><p><br/></p><p>在自己错的时候，更是要放低姿态，主动认错，而不是跟别人比谁的声音大、仗势高、嗓声亮。</p><p><br/></p><p>生活中总有一些人他们仗势欺人，随意乱发脾气。他们以为权势地位能得到别人的尊重和敬仰，其实那只是别人不愿揭穿你，表意附和你罢了。</p><p><br/></p><p>那些真正被人们深深地记在心里，且打心眼里尊重的人却是真正有修养的人。而一个有修养人都有一个好脾气，修养从来与物质毫无瓜葛。</p><p><br/></p><p>一个有修养的人，必须具备很多优秀的品质，但好脾气，绝对算得上是这所有品质中最重要的一种。</p></section></section></section></section><section><p><span>- END -</span></p>
							<section data-id="2484" data-tools="135编辑器"><section data-role="outer" label="Powered by 135editor.com"><section data-id="2484" data-tools="135编辑器">
							<section data-width="90%"><section data-style="font-size:14px;text-align:center"><fieldset><section><p><span>作者：李思圆，微信公众号：温暖的女子（ID：wennuan-312），新浪微博：@饮水-思圆，新书《生活需要仪式感》正在热销中。专栏作者，写安静从容，有温度，有力量的原创暖文。点击“<strong>阅读原文</strong>”即可在当当网购书。</span></p>
							<p><br/></p><p><img data-w="1220" data-copyright="0" data-ratio="1.2295081967213115" data-s="300,640" src="https://mmbiz.qpic.cn/mmbiz_jpg/KBQgh6ibe49mTOc2G5z2rt921drib19QL4XDS5vKowuZZicOViaggvGdfGaOahSQuRjMfWUx4U07OicTVh6fBzfmia1A/640?wx_fmt=jpeg" data-type="jpeg"/></p><p><br/></p></section></fieldset></section></section></section></section></section></section>
		                
				</div>
				

	</body></html>
	`,
		`<html><head></head><body>
	
		<div>   
		<p>
		<img data-ratio="0.09477521263669501" data-type="png" data-s="300,640" data-w="1646" src="https://mmbiz.qpic.cn/mmbiz_png/8xT0JYxK1STkibtsjBSq5oX6EuOhG1na3baJwpAOBWnLJUEVkpAegOibaKoibMibVuskQrnWXnDm7EiaUvjsv0iacdGw/0?wx_fmt=png"/></p>
		<p><br/></p>
		<p><span>我在马路边，捡到一分钱。怎么办？</span></p>
		<p><span>一个48秒的小魔术▼▼▼</span></p>
		<p><br/></p><p></p><p><br/></p><p>
		<span>开会、挤地铁、接孩子放学，月底流量不足又找不到wifi的话，</span>
		<span>就先看精简图片版吧。</span>
		<br/></p><p><br/></p><p>
		<img data-ratio="0.6208333333333333" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5P19MPViau5ADAuX7q5kMX0BWarTRDD8FicPc7gVReVNfiaYj0MVTmgmdw/0?wx_fmt=jpeg"/></p>
		<p><img data-ratio="0.5777777777777777" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5uxmbTiaQOXHOjwWvHzfMYv8SaElFbSRkXrflRz0d9D6iarh0ibYymey8w/0?wx_fmt=jpeg"/></p>
		<p><img data-ratio="0.575" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5Elv4G2fzAtN33dmvgNJxlzAxQvj5qWs2gdDHYX9z8ouu8vIfZEr79Q/0?wx_fmt=jpeg"/></p>
		<p><img data-ratio="0.5805555555555556" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe5fwabpvSZpBLdZhqQsSY0xiaOsJeUL3DgkvQrPceonj8Z5ARZKricPJww/0?wx_fmt=jpeg"/></p>
		<p><img data-ratio="0.5847222222222222" data-type="jpeg" data-s="300,640" data-w="720" src="https://mmbiz.qpic.cn/mmbiz_jpg/8xT0JYxK1SRXftJ27LnV6y1EmolWnMe50pz3jHdIB6lvu3TAibLCJ4xEbwibUSA41rFxTBMtsibwe6HU3pUXkjwxA/0?wx_fmt=jpeg"/></p>
		<p><br/></p>
		<p><span><strong><span>有奖互个动</span></strong></span></p>
		<p><span>可以直接刷二维码进出站的是哪条线？</span><span>评论区回答有点厉害的几位粉丝，发狗。</span></p>
		<p><span>*请注意！我台有奖互动截至时间为：每周五上午10点钟整，本台不支持刷赞哦。</span></p>
		<p><br/></p>
		<p><span>戳</span>
		<span><strong><span>阅读原文</span>
		</strong></span><span><strong><span>
		</span></strong></span>
		<span><br/>京东支付教你一分钱坐北京地铁。</span>
		</p><p><span><br/></span></p>
		<p>
		<img data-ratio="0.21113243761996161" data-type="png" data-s="300,640" data-w="1563" src="https://mmbiz.qpic.cn/mmbiz_png/8xT0JYxK1SQpcwChfZT3Ip3knKQqHR8GkNOM8fqeSEwkz2LoIqbctjplIP96ESyeAhmMKNgJuVSTmKq3RM958Q/0?wx_fmt=png"/></p><p><br/>
		</p>
				</div>

	</body></html>
	`}
)

func TestHtml2md(t *testing.T) {
	for _, hm := range htmlx {
		mk := Convert(hm)
		t.Fatal(mk)
	}
}
