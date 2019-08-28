<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#375EAB">

  <title>task2/animal/animal2.go - The Go Programming Language</title>

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
    <a href="/task2">task2</a>/<a href="/task2/animal">animal</a>/<span class="text-muted">animal2.go</span>
  </h1>





  <h2>
    Documentation: <a href="/pkg/task2/animal">task2/animal</a>
  </h2>



<div id="nav"></div>


<script type='text/javascript'>document.ANALYSIS_DATA = null;</script>
<pre><span id="L1" class="ln">     1&nbsp;&nbsp;</span>package main
<span id="L2" class="ln">     2&nbsp;&nbsp;</span>
<span id="L3" class="ln">     3&nbsp;&nbsp;</span>import (
<span id="L4" class="ln">     4&nbsp;&nbsp;</span>	&#34;fmt&#34;
<span id="L5" class="ln">     5&nbsp;&nbsp;</span>)
<span id="L6" class="ln">     6&nbsp;&nbsp;</span>
<span id="L7" class="ln">     7&nbsp;&nbsp;</span>type Animal interface{
<span id="L8" class="ln">     8&nbsp;&nbsp;</span>	Eat()
<span id="L9" class="ln">     9&nbsp;&nbsp;</span>	Move()
<span id="L10" class="ln">    10&nbsp;&nbsp;</span>	Speak()
<span id="L11" class="ln">    11&nbsp;&nbsp;</span>}
<span id="L12" class="ln">    12&nbsp;&nbsp;</span>
<span id="L13" class="ln">    13&nbsp;&nbsp;</span>type Cow struct {
<span id="L14" class="ln">    14&nbsp;&nbsp;</span>	food string
<span id="L15" class="ln">    15&nbsp;&nbsp;</span>	locomotion string
<span id="L16" class="ln">    16&nbsp;&nbsp;</span>	noise string
<span id="L17" class="ln">    17&nbsp;&nbsp;</span>}
<span id="L18" class="ln">    18&nbsp;&nbsp;</span>
<span id="L19" class="ln">    19&nbsp;&nbsp;</span>type Bird struct {
<span id="L20" class="ln">    20&nbsp;&nbsp;</span>	food string
<span id="L21" class="ln">    21&nbsp;&nbsp;</span>	locomotion string
<span id="L22" class="ln">    22&nbsp;&nbsp;</span>	noise string
<span id="L23" class="ln">    23&nbsp;&nbsp;</span>}
<span id="L24" class="ln">    24&nbsp;&nbsp;</span>
<span id="L25" class="ln">    25&nbsp;&nbsp;</span>type Snake struct {
<span id="L26" class="ln">    26&nbsp;&nbsp;</span>	food string
<span id="L27" class="ln">    27&nbsp;&nbsp;</span>	locomotion string
<span id="L28" class="ln">    28&nbsp;&nbsp;</span>	noise string
<span id="L29" class="ln">    29&nbsp;&nbsp;</span>}
<span id="L30" class="ln">    30&nbsp;&nbsp;</span>func (c Cow) Eat(){
<span id="L31" class="ln">    31&nbsp;&nbsp;</span>	fmt.Println(c.food)
<span id="L32" class="ln">    32&nbsp;&nbsp;</span>}
<span id="L33" class="ln">    33&nbsp;&nbsp;</span>func (b Bird) Eat(){
<span id="L34" class="ln">    34&nbsp;&nbsp;</span>        fmt.Println(b.food)
<span id="L35" class="ln">    35&nbsp;&nbsp;</span>}
<span id="L36" class="ln">    36&nbsp;&nbsp;</span>func (s Snake) Eat(){
<span id="L37" class="ln">    37&nbsp;&nbsp;</span>	fmt.Println(s.food)
<span id="L38" class="ln">    38&nbsp;&nbsp;</span>}
<span id="L39" class="ln">    39&nbsp;&nbsp;</span>func (c Cow) Move(){
<span id="L40" class="ln">    40&nbsp;&nbsp;</span>	fmt.Println(c.locomotion)
<span id="L41" class="ln">    41&nbsp;&nbsp;</span>}
<span id="L42" class="ln">    42&nbsp;&nbsp;</span>func (b Bird) Move(){
<span id="L43" class="ln">    43&nbsp;&nbsp;</span>	fmt.Println(b.locomotion)
<span id="L44" class="ln">    44&nbsp;&nbsp;</span>}
<span id="L45" class="ln">    45&nbsp;&nbsp;</span>func (s Snake) Move(){
<span id="L46" class="ln">    46&nbsp;&nbsp;</span>	fmt.Println(s.locomotion)
<span id="L47" class="ln">    47&nbsp;&nbsp;</span>}
<span id="L48" class="ln">    48&nbsp;&nbsp;</span>func (c Cow) Speak(){
<span id="L49" class="ln">    49&nbsp;&nbsp;</span>	fmt.Println(c.noise)
<span id="L50" class="ln">    50&nbsp;&nbsp;</span>}
<span id="L51" class="ln">    51&nbsp;&nbsp;</span>func (b Bird) Speak(){
<span id="L52" class="ln">    52&nbsp;&nbsp;</span>	fmt.Println(b.noise)
<span id="L53" class="ln">    53&nbsp;&nbsp;</span>}
<span id="L54" class="ln">    54&nbsp;&nbsp;</span>func (s Snake) Speak(){
<span id="L55" class="ln">    55&nbsp;&nbsp;</span>	fmt.Println(s.noise)
<span id="L56" class="ln">    56&nbsp;&nbsp;</span>}
<span id="L57" class="ln">    57&nbsp;&nbsp;</span>func Eat_inter(a Animal){
<span id="L58" class="ln">    58&nbsp;&nbsp;</span>	a.Eat()
<span id="L59" class="ln">    59&nbsp;&nbsp;</span>}
<span id="L60" class="ln">    60&nbsp;&nbsp;</span>func Move_inter(a Animal){
<span id="L61" class="ln">    61&nbsp;&nbsp;</span>	a.Move()
<span id="L62" class="ln">    62&nbsp;&nbsp;</span>}
<span id="L63" class="ln">    63&nbsp;&nbsp;</span>func Speak_inter(a Animal){
<span id="L64" class="ln">    64&nbsp;&nbsp;</span>	a.Speak()
<span id="L65" class="ln">    65&nbsp;&nbsp;</span>}
<span id="L66" class="ln">    66&nbsp;&nbsp;</span>
<span id="L67" class="ln">    67&nbsp;&nbsp;</span>func main() {
<span id="L68" class="ln">    68&nbsp;&nbsp;</span>	cow := Cow{&#34;grass&#34;,&#34;walk&#34;,&#34;moo&#34;}
<span id="L69" class="ln">    69&nbsp;&nbsp;</span>	bird := Bird{&#34;worms&#34;,&#34;fly&#34;,&#34;peep&#34;}
<span id="L70" class="ln">    70&nbsp;&nbsp;</span>	snake := Snake{&#34;mice&#34;,&#34;slither&#34;,&#34;hsss&#34;}
<span id="L71" class="ln">    71&nbsp;&nbsp;</span>	
<span id="L72" class="ln">    72&nbsp;&nbsp;</span>	var command, name, request string
<span id="L73" class="ln">    73&nbsp;&nbsp;</span>
<span id="L74" class="ln">    74&nbsp;&nbsp;</span>	animal := make(map[string]Animal)
<span id="L75" class="ln">    75&nbsp;&nbsp;</span>
<span id="L76" class="ln">    76&nbsp;&nbsp;</span>	for {
<span id="L77" class="ln">    77&nbsp;&nbsp;</span>		fmt.Println(&#34;&gt;&#34;)
<span id="L78" class="ln">    78&nbsp;&nbsp;</span>		fmt.Scanln(&amp;command,&amp;name,&amp;request)
<span id="L79" class="ln">    79&nbsp;&nbsp;</span>		if command == &#34;newanimal&#34;{
<span id="L80" class="ln">    80&nbsp;&nbsp;</span>			switch request{
<span id="L81" class="ln">    81&nbsp;&nbsp;</span>			case &#34;cow&#34;:
<span id="L82" class="ln">    82&nbsp;&nbsp;</span>				animal[name] = cow
<span id="L83" class="ln">    83&nbsp;&nbsp;</span>			case &#34;bird&#34;:
<span id="L84" class="ln">    84&nbsp;&nbsp;</span>				animal[name] = bird
<span id="L85" class="ln">    85&nbsp;&nbsp;</span>			case &#34;snake&#34;:
<span id="L86" class="ln">    86&nbsp;&nbsp;</span>				animal[name] = snake
<span id="L87" class="ln">    87&nbsp;&nbsp;</span>			}
<span id="L88" class="ln">    88&nbsp;&nbsp;</span>			fmt.Println(&#34;Created it!&#34;)
<span id="L89" class="ln">    89&nbsp;&nbsp;</span>		} else if command == &#34;query&#34;{
<span id="L90" class="ln">    90&nbsp;&nbsp;</span>			switch request{
<span id="L91" class="ln">    91&nbsp;&nbsp;</span>			case &#34;eat&#34;:
<span id="L92" class="ln">    92&nbsp;&nbsp;</span>				Eat_inter(animal[name])
<span id="L93" class="ln">    93&nbsp;&nbsp;</span>			case &#34;move&#34;:
<span id="L94" class="ln">    94&nbsp;&nbsp;</span>				Move_inter(animal[name])
<span id="L95" class="ln">    95&nbsp;&nbsp;</span>			case &#34;speak&#34;:
<span id="L96" class="ln">    96&nbsp;&nbsp;</span>				Speak_inter(animal[name])
<span id="L97" class="ln">    97&nbsp;&nbsp;</span>			}
<span id="L98" class="ln">    98&nbsp;&nbsp;</span>		} else{
<span id="L99" class="ln">    99&nbsp;&nbsp;</span>			fmt.Println(&#34;Unknow command!!&#34;)
<span id="L100" class="ln">   100&nbsp;&nbsp;</span>		}
<span id="L101" class="ln">   101&nbsp;&nbsp;</span>	}
<span id="L102" class="ln">   102&nbsp;&nbsp;</span>}
<span id="L103" class="ln">   103&nbsp;&nbsp;</span>
</pre><p><a href="/task2/animal/animal2.go?m=text">View as plain text</a></p>

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

