<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#375EAB">

  <title>task2/displacement/displacement.go - The Go Programming Language</title>

<link type="text/css" rel="stylesheet" href="/lib/godoc/style.css">


<script>window.initFuncs = [];</script>

<script src="/lib/godoc/jquery.js" defer></script>



<script>var goVersion = "go1.12.4";</script>
<script src="/lib/godoc/godocs.js" defer></script>
</head>
<body>

<div id='lowframe' style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
...
</div><!-- #lowframe -->

<div id="topbar" class="wide"><div class="container">
<div class="top-heading" id="heading-wide"><a href="/">The Go Programming Language</a></div>
<div class="top-heading" id="heading-narrow"><a href="/">Go</a></div>
<a href="#" id="menu-button"><span id="menu-button-arrow">&#9661;</span></a>
<form method="GET" action="/search">
<div id="menu">
<a href="/doc/">Documents</a>
<a href="/pkg/">Packages</a>
<a href="/project/">The Project</a>
<a href="/help/">Help</a>

<a href="/blog/">Blog</a>


<span class="search-box"><input type="search" id="search" name="q" placeholder="Search" aria-label="Search" required><button type="submit"><span><!-- magnifying glass: --><svg width="24" height="24" viewBox="0 0 24 24"><title>submit search</title><path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/><path d="M0 0h24v24H0z" fill="none"/></svg></span></button></span>
</div>
</form>

</div></div>



<div id="page" class="wide">
<div class="container">


  <h1>
    Source file
    <a href="/task2">task2</a>/<a href="/task2/displacement">displacement</a>/<span class="text-muted">displacement.go</span>
  </h1>





  <h2>
    Documentation: <a href="/pkg/task2/displacement">task2/displacement</a>
  </h2>



<div id="nav"></div>


<script type='text/javascript'>document.ANALYSIS_DATA = null;</script>
<pre><span id="L1" class="ln">     1&nbsp;&nbsp;</span>package main
<span id="L2" class="ln">     2&nbsp;&nbsp;</span>
<span id="L3" class="ln">     3&nbsp;&nbsp;</span>import (
<span id="L4" class="ln">     4&nbsp;&nbsp;</span>	&#34;fmt&#34;
<span id="L5" class="ln">     5&nbsp;&nbsp;</span>	&#34;math&#34;
<span id="L6" class="ln">     6&nbsp;&nbsp;</span>	&#34;time&#34;
<span id="L7" class="ln">     7&nbsp;&nbsp;</span>)
<span id="L8" class="ln">     8&nbsp;&nbsp;</span>
<span id="L9" class="ln">     9&nbsp;&nbsp;</span>func GenDisplaceFn(acc, vel, dis float64) func (float64) float64{
<span id="L10" class="ln">    10&nbsp;&nbsp;</span>	fn := func(time float64) float64{
<span id="L11" class="ln">    11&nbsp;&nbsp;</span>		return (0.5)*(acc)*(math.Pow(time,2)) + vel*time + dis
<span id="L12" class="ln">    12&nbsp;&nbsp;</span>	} 
<span id="L13" class="ln">    13&nbsp;&nbsp;</span>	return fn
<span id="L14" class="ln">    14&nbsp;&nbsp;</span>}
<span id="L15" class="ln">    15&nbsp;&nbsp;</span>
<span id="L16" class="ln">    16&nbsp;&nbsp;</span>func main(){
<span id="L17" class="ln">    17&nbsp;&nbsp;</span>	fmt.Print(&#34;Enter the acceleration: &#34;)
<span id="L18" class="ln">    18&nbsp;&nbsp;</span>	var acce float64
<span id="L19" class="ln">    19&nbsp;&nbsp;</span>	fmt.Scan(&amp;acce)
<span id="L20" class="ln">    20&nbsp;&nbsp;</span>
<span id="L21" class="ln">    21&nbsp;&nbsp;</span>	fmt.Print(&#34;Enter the initial velocity: &#34;)
<span id="L22" class="ln">    22&nbsp;&nbsp;</span>	var vel float64
<span id="L23" class="ln">    23&nbsp;&nbsp;</span>	fmt.Scan(&amp;vel)
<span id="L24" class="ln">    24&nbsp;&nbsp;</span>
<span id="L25" class="ln">    25&nbsp;&nbsp;</span>	fmt.Print(&#34;Enter the initial displacement: &#34;)
<span id="L26" class="ln">    26&nbsp;&nbsp;</span>	var dis float64
<span id="L27" class="ln">    27&nbsp;&nbsp;</span>	fmt.Scan(&amp;dis)
<span id="L28" class="ln">    28&nbsp;&nbsp;</span>
<span id="L29" class="ln">    29&nbsp;&nbsp;</span>	fmt.Print(&#34;Enter the time: &#34;)
<span id="L30" class="ln">    30&nbsp;&nbsp;</span>	var tim float64
<span id="L31" class="ln">    31&nbsp;&nbsp;</span>	fmt.Scan(&amp;tim)
<span id="L32" class="ln">    32&nbsp;&nbsp;</span>
<span id="L33" class="ln">    33&nbsp;&nbsp;</span>	time.Sleep(time.Duration(tim) * time.Second)
<span id="L34" class="ln">    34&nbsp;&nbsp;</span>
<span id="L35" class="ln">    35&nbsp;&nbsp;</span>	Dist := GenDisplaceFn(acce,vel,dis)
<span id="L36" class="ln">    36&nbsp;&nbsp;</span>	fmt.Println(Dist(tim))
<span id="L37" class="ln">    37&nbsp;&nbsp;</span>}
<span id="L38" class="ln">    38&nbsp;&nbsp;</span>
</pre><p><a href="/task2/displacement/displacement.go?m=text">View as plain text</a></p>

<div id="footer">
Build version go1.12.4.<br>
Except as <a href="https://developers.google.com/site-policies#restrictions">noted</a>,
the content of this page is licensed under the
Creative Commons Attribution 3.0 License,
and code is licensed under a <a href="/LICENSE">BSD license</a>.<br>
<a href="/doc/tos.html">Terms of Service</a> |
<a href="http://www.google.com/intl/en/policies/privacy/">Privacy Policy</a>
</div>

</div><!-- .container -->
</div><!-- #page -->

</body>
</html>

