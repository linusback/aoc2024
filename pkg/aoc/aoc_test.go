package aoc

import (
	"bytes"
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"regexp"
	"strings"
	"testing"
)

func Benchmark_StandardConvert(b *testing.B) {
	b.SetBytes(14000)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(testHtml)
		_, err := htmltomarkdown.ConvertReader(r)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_RegexConvert(b *testing.B) {
	b.SetBytes(14000)
	b.ReportAllocs()
	reg, err := regexp.Compile("(?i)(?s)<main>(?P<main>.*)</main>")
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		by := reg.Find(testHtmlBytes)
		_, err = htmltomarkdown.ConvertReader(bytes.NewReader(by))
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_ConvertOnlyMain(b *testing.B) {
	b.SetBytes(5677)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := htmltomarkdown.ConvertReader(bytes.NewReader(testMainBytes))
		if err != nil {
			b.Error(err)
		}
	}
}

//func Benchmark_CusomReaderConvert(b *testing.B) {
//	b.SetBytes(14000)
//	b.ReportAllocs()
//	reg, err := regexp.Compile("(?i)(?s)<main>(?P<main>.*)</main>")
//	if err != nil {
//		b.Error(err)
//	}
//	for i := 0; i < b.N; i++ {
//		by := reg.Find(testHtmlBytes)
//		_, err = htmltomarkdown.ConvertReader(bytes.NewReader(by))
//		if err != nil {
//			b.Error(err)
//		}
//	}
//}

func Test_RegexConvert(t *testing.T) {
	reg, err := regexp.Compile("(?i)(?s)<main>(?P<main>.*)</main>")
	if err != nil {
		t.Error(err)
	}
	by := reg.Find(testHtmlBytes)
	_, err = htmltomarkdown.ConvertReader(bytes.NewReader(by))
	if err != nil {
		t.Error(err)
	}
	t.Log(len(testMainBytes))
	t.Log(string(by))
}

var (
	testHtmlBytes = []byte(testHtml)
	testMainBytes = []byte(`<article class="day-desc"><h2>--- Day 1: Historian Hysteria ---</h2><p>The <em>Chief Historian</em> is always present for the big Christmas sleigh launch, but nobody has seen him in months! Last anyone heard, he was visiting locations that are historically significant to the North Pole; a group of Senior Historians has asked you to accompany them as they check the places they think he was most likely to visit.</p>
<p>As each location is checked, they will mark it on their list with a <em class="star">star</em>. They figure the Chief Historian <em>must</em> be in one of the first fifty places they'll look, so in order to save Christmas, you need to help them get <em class="star">fifty stars</em> on their list before Santa takes off on December 25th.</p>
<p>Collect stars by solving puzzles.  Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first.  Each puzzle grants <em class="star">one star</em>. Good luck!</p>
<p>You haven't even left yet and the group of Elvish Senior Historians has already hit a problem: their list of locations to check is currently <em>empty</em>. Eventually, someone decides that the best place to check first would be the Chief Historian's office.</p>
<p>Upon pouring into the office, everyone confirms that the Chief Historian is indeed nowhere to be found. Instead, the Elves discover an assortment of notes and lists of historically significant locations! This seems to be the planning the Chief Historian was doing before he left. Perhaps these notes can be used to determine which locations to search?</p>
<p>Throughout the Chief's office, the historically significant locations are listed not by name but by a unique number called the <em>location ID</em>. To make sure they don't miss anything, The Historians split into two groups, each searching the office and trying to create their own complete list of location IDs.</p>
<p>There's just one problem: by holding the two lists up <em>side by side</em> (your puzzle input), it quickly becomes clear that the lists aren't very similar. Maybe you can help The Historians reconcile their lists?</p>
<p>For example:</p>
<pre><code>3   4
4   3
2   5
1   3
3   9
3   3
</code></pre>
<p>Maybe the lists are only off by a small amount! To find out, pair up the numbers and measure how far apart they are. Pair up the <em>smallest number in the left list</em> with the <em>smallest number in the right list</em>, then the <em>second-smallest left number</em> with the <em>second-smallest right number</em>, and so on.</p>
<p>Within each pair, figure out <em>how far apart</em> the two numbers are; you'll need to <em>add up all of those distances</em>. For example, if you pair up a <code>3</code> from the left list with a <code>7</code> from the right list, the distance apart is <code>4</code>; if you pair up a <code>9</code> with a <code>3</code>, the distance apart is <code>6</code>.</p>
<p>In the example list above, the pairs and distances would be as follows:</p>
<ul>
<li>The smallest number in the left list is <code>1</code>, and the smallest number in the right list is <code>3</code>. The distance between them is <code><em>2</em></code>.</li>
<li>The second-smallest number in the left list is <code>2</code>, and the second-smallest number in the right list is another <code>3</code>. The distance between them is <code><em>1</em></code>.</li>
<li>The third-smallest number in both lists is <code>3</code>, so the distance between them is <code><em>0</em></code>.</li>
<li>The next numbers to pair up are <code>3</code> and <code>4</code>, a distance of <code><em>1</em></code>.</li>
<li>The fifth-smallest numbers in each list are <code>3</code> and <code>5</code>, a distance of <code><em>2</em></code>.</li>
<li>Finally, the largest number in the left list is <code>4</code>, while the largest number in the right list is <code>9</code>; these are a distance <code><em>5</em></code> apart.</li>
</ul>
<p>To find the <em>total distance</em> between the left list and the right list, add up the distances between all of the pairs you found. In the example above, this is <code>2 + 1 + 0 + 1 + 2 + 5</code>, a total distance of <code><em>11</em></code>!</p>
<p>Your actual left and right lists contain many location IDs. <em>What is the total distance between your lists?</em></p>
</article>
<p>To begin, <a href="1/input" target="_blank">get your puzzle input</a>.</p>
<form method="post" action="1/answer"><input type="hidden" name="level" value="1"/><p>Answer: <input type="text" name="answer" autocomplete="off"/> <input type="submit" value="[Submit]"/></p></form>
<p>You can also <span class="share">[Share<span class="share-content">on
  <a href="https://bsky.app/intent/compose?text=%22Historian+Hysteria%22+%2D+Day+1+%2D+Advent+of+Code+2024+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2024%2Fday%2F1" target="_blank">Bluesky</a>
  <a href="https://twitter.com/intent/tweet?text=%22Historian+Hysteria%22+%2D+Day+1+%2D+Advent+of+Code+2024&amp;url=https%3A%2F%2Fadventofcode%2Ecom%2F2024%2Fday%2F1&amp;related=ericwastl&amp;hashtags=AdventOfCode" target="_blank">Twitter</a>
  <a href="javascript:void(0);" onclick="var ms; try{ms=localStorage.getItem('mastodon.server')}finally{} if(typeof ms!=='string')ms=''; ms=prompt('Mastodon Server?',ms); if(typeof ms==='string' && ms.length){this.href='https://'+ms+'/share?text=%22Historian+Hysteria%22+%2D+Day+1+%2D+Advent+of+Code+2024+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2024%2Fday%2F1';try{localStorage.setItem('mastodon.server',ms);}finally{}}else{return false;}" target="_blank">Mastodon</a
></span>]</span> this puzzle.</p>`)
)

const testHtml = `<!DOCTYPE html>
<html lang="en-us">
<head>
<meta charset="utf-8"/>
<title>Day 1 - Advent of Code 2024</title>
<link rel="stylesheet" type="text/css" href="/static/style.css?31"/>
<link rel="stylesheet alternate" type="text/css" href="/static/highcontrast.css?1" title="High Contrast"/>
<link rel="shortcut icon" href="/favicon.png"/>
<script>window.addEventListener('click', function(e,s,r){if(e.target.nodeName==='CODE'&&e.detail===3){s=window.getSelection();s.removeAllRanges();r=document.createRange();r.selectNodeContents(e.target);s.addRange(r);}});</script>
</head><!--




Oh, hello!  Funny seeing you here.

I appreciate your enthusiasm, but you aren't going to find much down here.
There certainly aren't clues to any of the puzzles.  The best surprises don't
even appear in the source until you unlock them for real.

Please be careful with automated requests; I'm not a massive company, and I can
only take so much traffic.  Please be considerate so that everyone gets to play.

If you're curious about how Advent of Code works, it's running on some custom
Perl code. Other than a few integrations (auth, analytics, social media), I
built the whole thing myself, including the design, animations, prose, and all
of the puzzles.

The puzzles are most of the work; preparing a new calendar and a new set of
puzzles each year takes all of my free time for 4-5 months. A lot of effort
went into building this thing - I hope you're enjoying playing it as much as I
enjoyed making it for you!

If you'd like to hang out, I'm @was.tl on Bluesky, @ericwastl@hachyderm.io on
Mastodon, and @ericwastl on Twitter.

- Eric Wastl


















































-->
<body>
<header><div><h1 class="title-global"><a href="/">Advent of Code</a></h1><nav><ul><li><a href="/2024/about">[About]</a></li><li><a href="/2024/events">[Events]</a></li><li><a href="https://cottonbureau.com/people/advent-of-code" target="_blank">[Shop]</a></li><li><a href="/2024/settings">[Settings]</a></li><li><a href="/2024/auth/logout">[Log Out]</a></li></ul></nav><div class="user">Linus A Back</div></div><div><h1 class="title-event">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span class="title-event-wrap">y(</span><a href="/2024">2024</a><span class="title-event-wrap">)</span></h1><nav><ul><li><a href="/2024">[Calendar]</a></li><li><a href="/2024/support">[AoC++]</a></li><li><a href="/2024/sponsors">[Sponsors]</a></li><li><a href="/2024/leaderboard">[Leaderboard]</a></li><li><a href="/2024/stats">[Stats]</a></li></ul></nav></div></header>

<div id="sidebar">
<div id="sponsor"><div class="quiet">Our <a href="/2024/sponsors">sponsors</a> help make Advent of Code possible:</div><div class="sponsor"><a href="/2024/sponsors/redirect?url=https%3A%2F%2Fcloudsmith%2Ecom%3Futm%5Fsource%3Daoc%26utm%5Fmedium%3Dsponsor%26utm%5Fcampaign%3Daoc2024" target="_blank" onclick="if(ga)ga('send','event','sponsor','sidebar',this.href);" rel="noopener">Cloudsmith</a> - Code shapes reality; now solve your next puzzle: global artifact management. Join our journey to shape, secure, and be the world&apos;s software supply chain. We overcame the impossible and became the improbable - now, it&apos;s inevitable.</div></div>
</div><!--/sidebar-->

<main>
<article class="day-desc"><h2>--- Day 1: Historian Hysteria ---</h2><p>The <em>Chief Historian</em> is always present for the big Christmas sleigh launch, but nobody has seen him in months! Last anyone heard, he was visiting locations that are historically significant to the North Pole; a group of Senior Historians has asked you to accompany them as they check the places they think he was most likely to visit.</p>
<p>As each location is checked, they will mark it on their list with a <em class="star">star</em>. They figure the Chief Historian <em>must</em> be in one of the first fifty places they'll look, so in order to save Christmas, you need to help them get <em class="star">fifty stars</em> on their list before Santa takes off on December 25th.</p>
<p>Collect stars by solving puzzles.  Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first.  Each puzzle grants <em class="star">one star</em>. Good luck!</p>
<p>You haven't even left yet and the group of Elvish Senior Historians has already hit a problem: their list of locations to check is currently <em>empty</em>. Eventually, someone decides that the best place to check first would be the Chief Historian's office.</p>
<p>Upon pouring into the office, everyone confirms that the Chief Historian is indeed nowhere to be found. Instead, the Elves discover an assortment of notes and lists of historically significant locations! This seems to be the planning the Chief Historian was doing before he left. Perhaps these notes can be used to determine which locations to search?</p>
<p>Throughout the Chief's office, the historically significant locations are listed not by name but by a unique number called the <em>location ID</em>. To make sure they don't miss anything, The Historians split into two groups, each searching the office and trying to create their own complete list of location IDs.</p>
<p>There's just one problem: by holding the two lists up <em>side by side</em> (your puzzle input), it quickly becomes clear that the lists aren't very similar. Maybe you can help The Historians reconcile their lists?</p>
<p>For example:</p>
<pre><code>3   4
4   3
2   5
1   3
3   9
3   3
</code></pre>
<p>Maybe the lists are only off by a small amount! To find out, pair up the numbers and measure how far apart they are. Pair up the <em>smallest number in the left list</em> with the <em>smallest number in the right list</em>, then the <em>second-smallest left number</em> with the <em>second-smallest right number</em>, and so on.</p>
<p>Within each pair, figure out <em>how far apart</em> the two numbers are; you'll need to <em>add up all of those distances</em>. For example, if you pair up a <code>3</code> from the left list with a <code>7</code> from the right list, the distance apart is <code>4</code>; if you pair up a <code>9</code> with a <code>3</code>, the distance apart is <code>6</code>.</p>
<p>In the example list above, the pairs and distances would be as follows:</p>
<ul>
<li>The smallest number in the left list is <code>1</code>, and the smallest number in the right list is <code>3</code>. The distance between them is <code><em>2</em></code>.</li>
<li>The second-smallest number in the left list is <code>2</code>, and the second-smallest number in the right list is another <code>3</code>. The distance between them is <code><em>1</em></code>.</li>
<li>The third-smallest number in both lists is <code>3</code>, so the distance between them is <code><em>0</em></code>.</li>
<li>The next numbers to pair up are <code>3</code> and <code>4</code>, a distance of <code><em>1</em></code>.</li>
<li>The fifth-smallest numbers in each list are <code>3</code> and <code>5</code>, a distance of <code><em>2</em></code>.</li>
<li>Finally, the largest number in the left list is <code>4</code>, while the largest number in the right list is <code>9</code>; these are a distance <code><em>5</em></code> apart.</li>
</ul>
<p>To find the <em>total distance</em> between the left list and the right list, add up the distances between all of the pairs you found. In the example above, this is <code>2 + 1 + 0 + 1 + 2 + 5</code>, a total distance of <code><em>11</em></code>!</p>
<p>Your actual left and right lists contain many location IDs. <em>What is the total distance between your lists?</em></p>
</article>
<p>To begin, <a href="1/input" target="_blank">get your puzzle input</a>.</p>
<form method="post" action="1/answer"><input type="hidden" name="level" value="1"/><p>Answer: <input type="text" name="answer" autocomplete="off"/> <input type="submit" value="[Submit]"/></p></form>
<p>You can also <span class="share">[Share<span class="share-content">on
  <a href="https://bsky.app/intent/compose?text=%22Historian+Hysteria%22+%2D+Day+1+%2D+Advent+of+Code+2024+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2024%2Fday%2F1" target="_blank">Bluesky</a>
  <a href="https://twitter.com/intent/tweet?text=%22Historian+Hysteria%22+%2D+Day+1+%2D+Advent+of+Code+2024&amp;url=https%3A%2F%2Fadventofcode%2Ecom%2F2024%2Fday%2F1&amp;related=ericwastl&amp;hashtags=AdventOfCode" target="_blank">Twitter</a>
  <a href="javascript:void(0);" onclick="var ms; try{ms=localStorage.getItem('mastodon.server')}finally{} if(typeof ms!=='string')ms=''; ms=prompt('Mastodon Server?',ms); if(typeof ms==='string' && ms.length){this.href='https://'+ms+'/share?text=%22Historian+Hysteria%22+%2D+Day+1+%2D+Advent+of+Code+2024+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2024%2Fday%2F1';try{localStorage.setItem('mastodon.server',ms);}finally{}}else{return false;}" target="_blank">Mastodon</a
></span>]</span> this puzzle.</p>
</main>

<!-- ga -->
<script>
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
ga('create', 'UA-69522494-1', 'auto');
ga('set', 'anonymizeIp', true);
ga('send', 'pageview');
</script>
<!-- /ga -->
</body>
</html>`