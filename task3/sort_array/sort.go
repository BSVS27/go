<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#375EAB">

  <title>task3/sort_array/sort.go - The Go Programming Language</title>

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
    <a href="/task3">task3</a>/<a href="/task3/sort_array">sort_array</a>/<span class="text-muted">sort.go</span>
  </h1>





  <h2>
    Documentation: <a href="/pkg/task3/sort_array">task3/sort_array</a>
  </h2>



<div id="nav"></div>


<script type='text/javascript'>document.ANALYSIS_DATA = null;</script>
<pre><span id="L1" class="ln">     1&nbsp;&nbsp;</span>package main
<span id="L2" class="ln">     2&nbsp;&nbsp;</span>
<span id="L3" class="ln">     3&nbsp;&nbsp;</span>import (
<span id="L4" class="ln">     4&nbsp;&nbsp;</span>	&#34;fmt&#34;
<span id="L5" class="ln">     5&nbsp;&nbsp;</span>	&#34;math&#34;
<span id="L6" class="ln">     6&nbsp;&nbsp;</span>)
<span id="L7" class="ln">     7&nbsp;&nbsp;</span>
<span id="L8" class="ln">     8&nbsp;&nbsp;</span>func sort(s []int, c chan []int){
<span id="L9" class="ln">     9&nbsp;&nbsp;</span>	var aux int
<span id="L10" class="ln">    10&nbsp;&nbsp;</span>	for i:=0;i&lt;(len(s)-1);i++{
<span id="L11" class="ln">    11&nbsp;&nbsp;</span>		for i:=0;i&lt;(len(s)-1);i++{
<span id="L12" class="ln">    12&nbsp;&nbsp;</span>			if s[i] &gt; s[i+1]{
<span id="L13" class="ln">    13&nbsp;&nbsp;</span>				aux = s[i]
<span id="L14" class="ln">    14&nbsp;&nbsp;</span>				s[i] = s[i+1]
<span id="L15" class="ln">    15&nbsp;&nbsp;</span>				s[i+1] = aux
<span id="L16" class="ln">    16&nbsp;&nbsp;</span>			}
<span id="L17" class="ln">    17&nbsp;&nbsp;</span>		}
<span id="L18" class="ln">    18&nbsp;&nbsp;</span>	}
<span id="L19" class="ln">    19&nbsp;&nbsp;</span>	fmt.Println(s)
<span id="L20" class="ln">    20&nbsp;&nbsp;</span>	c &lt;- s
<span id="L21" class="ln">    21&nbsp;&nbsp;</span>}
<span id="L22" class="ln">    22&nbsp;&nbsp;</span>
<span id="L23" class="ln">    23&nbsp;&nbsp;</span>func main() {
<span id="L24" class="ln">    24&nbsp;&nbsp;</span>	fmt.Println(&#34;How many integers do you want to enter?&#34;)
<span id="L25" class="ln">    25&nbsp;&nbsp;</span>	var in int
<span id="L26" class="ln">    26&nbsp;&nbsp;</span>	fmt.Scan(&amp;in)
<span id="L27" class="ln">    27&nbsp;&nbsp;</span>	array := make([]int,0)
<span id="L28" class="ln">    28&nbsp;&nbsp;</span>	for i:=0;i&lt;in;i++{
<span id="L29" class="ln">    29&nbsp;&nbsp;</span>		fmt.Println(&#34;Please enter the value &#34;,i,&#34; of the array: &#34;)
<span id="L30" class="ln">    30&nbsp;&nbsp;</span>		var val int
<span id="L31" class="ln">    31&nbsp;&nbsp;</span>		fmt.Scan(&amp;val)
<span id="L32" class="ln">    32&nbsp;&nbsp;</span>		array = append(array, val)
<span id="L33" class="ln">    33&nbsp;&nbsp;</span>	}
<span id="L34" class="ln">    34&nbsp;&nbsp;</span>	fmt.Println(&#34;The initial array is: &#34;,array)
<span id="L35" class="ln">    35&nbsp;&nbsp;</span>	chunk := int(math.Round(float64(in)/float64(4)))
<span id="L36" class="ln">    36&nbsp;&nbsp;</span>	c := make(chan []int)
<span id="L37" class="ln">    37&nbsp;&nbsp;</span>	go sort(array[:chunk],c)
<span id="L38" class="ln">    38&nbsp;&nbsp;</span>	go sort(array[chunk:2*chunk],c)
<span id="L39" class="ln">    39&nbsp;&nbsp;</span>	go sort(array[2*chunk:3*chunk],c)
<span id="L40" class="ln">    40&nbsp;&nbsp;</span>	go sort(array[3*chunk:],c)
<span id="L41" class="ln">    41&nbsp;&nbsp;</span>	a1 := &lt;-c
<span id="L42" class="ln">    42&nbsp;&nbsp;</span>	b1 := &lt;-c
<span id="L43" class="ln">    43&nbsp;&nbsp;</span>	c1 := &lt;-c
<span id="L44" class="ln">    44&nbsp;&nbsp;</span>	d1 := &lt;-c
<span id="L45" class="ln">    45&nbsp;&nbsp;</span>	s_array := make([]int,0)
<span id="L46" class="ln">    46&nbsp;&nbsp;</span>	for i:=0;i&lt;len(a1);i++{
<span id="L47" class="ln">    47&nbsp;&nbsp;</span>		s_array = append(s_array,a1[i])
<span id="L48" class="ln">    48&nbsp;&nbsp;</span>	}
<span id="L49" class="ln">    49&nbsp;&nbsp;</span>	for i:=0;i&lt;len(b1);i++{
<span id="L50" class="ln">    50&nbsp;&nbsp;</span>		s_array = append(s_array,b1[i])
<span id="L51" class="ln">    51&nbsp;&nbsp;</span>	}
<span id="L52" class="ln">    52&nbsp;&nbsp;</span>	for i:=0;i&lt;len(c1);i++{
<span id="L53" class="ln">    53&nbsp;&nbsp;</span>		s_array = append(s_array,c1[i])
<span id="L54" class="ln">    54&nbsp;&nbsp;</span>	}
<span id="L55" class="ln">    55&nbsp;&nbsp;</span>	for i:=0;i&lt;len(d1);i++{
<span id="L56" class="ln">    56&nbsp;&nbsp;</span>		s_array = append(s_array,d1[i])
<span id="L57" class="ln">    57&nbsp;&nbsp;</span>	}
<span id="L58" class="ln">    58&nbsp;&nbsp;</span>	fmt.Println(&#34;The sorted array is: &#34;,s_array)
<span id="L59" class="ln">    59&nbsp;&nbsp;</span>	fmt.Println(&#34;Done&#34;)
<span id="L60" class="ln">    60&nbsp;&nbsp;</span>}
<span id="L61" class="ln">    61&nbsp;&nbsp;</span>
</pre><p><a href="/task3/sort_array/sort.go?m=text">View as plain text</a></p>

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

