package regexps

import (
	"regexp"
	"strings"
    "unicode"
)

func SpaceMap(str string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsSpace(r) {
            return -1
        }
        return r
    }, str)
}


//var href = regexp.MustCompile(`<a href="(http:)?\/\/www\.bigmir\.net\/" target="_blank".+><\/a>`)
//var href = regexp.MustCompile(`<a href="(http:)?\/\/www\.bigmir\.net\/".+>`)
var href = regexp.MustCompile(`<a href="(http:)?\/\/(www\.)?(top\.)?bigmir\.net\/?(stat\/\d+)?".+>`)

var counterRxFull1 = regexp.MustCompile(`<!--bigmir\)net TOP 100-->
<script type="text\/javascript" language="javascript"><!--
bmN=navigator,bmD=document,bmD\.cookie='b=b',i=0,bs=\[\],bm={v:\d+,s:\d+,t:\d+,c:bmD\.cookie\?1:0,n:Math\.round\(\(Math\.random\(\)\* 1000000\)\),w:0};
for\(var f=self;f!=f\.parent;f=f\.parent\)bm\.w\+\+;
try{if\(bmN\.plugins&&bmN\.mimeTypes\.length&&\(x=bmN\.plugins\['Shockwave Flash']\)\)bm\.m=parseInt\(x\.description\.replace\(\/\(\[a-zA-Z]\|\\s\)\+\/,''\)\);
else for\(var f=3;f<20;f\+\+\)if\(eval\('new ActiveXObject\("ShockwaveFlash\.ShockwaveFlash\.'\+f\+'"\)'\)\)bm\.m=f}catch\(e\){;}
try{bm\.y=bmN\.javaEnabled\(\)\?1:0}catch\(e\){;}
try{bmS=screen;bm\.v\^=bm\.d=bmS\.colorDepth\|\|bmS\.pixelDepth;bm\.v\^=bm\.r=bmS\.width}catch\(e\){;}
r=bmD\.referrer\.replace\(\/\^w\+:\\\/\\\/\/,''\);if\(r&&r\.split\('\/'\)\[0]!=window\.location\.host\){bm\.f=escape\(r\)\.slice\(0,400\);bm\.v\^=r\.length}
bm\.v\^=window\.location\.href\.length;for\(var x in bm\) if\(\/\^\[vstcnwmydrf]\$\/\.test\(x\)\) bs\[i\+\+]=x\+bm\[x];
bmD\.write\('<a href="http:\/\/www\.bigmir\.net\/" target="_blank" onClick="img=new Image\(\);img\.src="\/\/www\.bigmir\.net\/\?cl=\d+";"><img src="\/\/c\.bigmir\.net\/\?'\+bs\.join\('&'\)\+'"  width="\d+" height="\d+" border="0" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100"><\/a>'\);
\/\/-->
<\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t\d+" width="\d+" height="\d+" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>
<!--bigmir\)net TOP 100-->`)

var counterRxFull2 = regexp.MustCompile(`<!--bigmir\)net TOP 100-->
<script type="text\/javascript" language="javascript"><!--
function BM_Draw\(oBM_STAT\){
document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="margin:0;padding:0;font:10px Tahoma;width:100px;height:70px;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/curl\/b\d+\.gif\\'\);background-repeat:no-repeat;"><div style="height:44px;"><\/div><div style="padding-left:33px; padding-right: 2px;"><div style="float:left;color:#400000;">хосты<\/div><div style="float:right;color:#400000;">'\+oBM_STAT\.hosts\+'<\/div><br clear="all" \/><\/div><div style="padding-left:25px;padding-right: 2px;"><div style="float:left;color:#400000;">хиты<\/div><div style="float:right;color:#400000;">'\+oBM_STAT\.hits\+'<\/div><br clear="all" \/><\/div><\/div><\/td><\/tr><\/table>'\);
}
\/\/-->
<\/script>
<script type="text\/javascript" language="javascript"><!--
bmN=navigator,bmD=document,bmD\.cookie='b=b',i=0,bs=\[],bm={o:1,v:\d+,s:\d+,t:0,c:bmD\.cookie\?1:0,n:Math\.round\(\(Math\.random\(\)\* 1000000\)\),w:0};
for\(var f=self;f!=f\.parent;f=f\.parent\)bm\.w\+\+;
try{if\(bmN\.plugins&&bmN\.mimeTypes\.length&&\(x=bmN\.plugins\['Shockwave Flash']\)\)bm\.m=parseInt\(x\.description\.replace\(\/\(\[a-zA-Z]\|\\s\)\+\/,''\)\);
else for\(var f=3;f<20;f\+\+\)if\(eval\('new ActiveXObject\("ShockwaveFlash\.ShockwaveFlash\.'\+f\+'"\)'\)\)bm\.m=f}catch\(e\){;}
try{bm\.y=bmN\.javaEnabled\(\)\?1:0}catch\(e\){;}
try{bmS=screen;bm\.v\^=bm\.d=bmS\.colorDepth\|\|bmS\.pixelDepth;bm\.v\^=bm\.r=bmS\.width}catch\(e\){;}
r=bmD\.referrer\.replace\(\/\^w\+:\\\/\\\/\/,''\);if\(r&&r\.split\('\/'\)\[0]!=window\.location\.host\){bm\.f=escape\(r\)\.slice\(0,400\);bm\.v\^=r\.length}
bm\.v\^=window\.location\.href\.length;for\(var x in bm\) if\(\/\^\[ovstcnwmydrf]\$\/\.test\(x\)\) bs\[i\+\+]=x\+bm\[x];
bmD\.write\('<sc'\+'ript type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?'\+bs\.join\('&'\)\+'"><\/sc'\+'ript>'\);
\/\/-->
<\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t2" width="88" height="31" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>
<!--bigmir\)net TOP 100-->`)

var counterRxFull3 = regexp.MustCompile(`<!--bigmir\)net TOP 100-->
<script type="text\/javascript" language="javascript"><!--
function BM_Draw\(oBM_STAT\){
document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="margin:0px;padding:0px;font-size:1px;width:88px;"><div style="background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_top\.gif\\'\) no-repeat bottom;"> <\/div><div style="font:10px Tahoma;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_center\.gif\\'\);"><div style="text-align:center;"><a href="http:\/\/www\.bigmir\.net\/" target="_blank" style="color:#0000ab;text-decoration:none;font:10px Tahoma;">bigmir<span style="color:#ff0000;">\)<\/span>net<\/a><\/div><div style="margin-top:3px;padding: 0px 6px 0px 6px;color:#[a-z\d]+;"><div style="float:left;font:10px Tahoma;">'\+oBM_STAT\.hosts\+'<\/div><div style="float:right;font:10px Tahoma;">'\+oBM_STAT\.hits\+'<\/div><\/div><br clear="all"\/><\/div><div style="background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_bottom\.gif\\'\) no-repeat top;"> <\/div><\/div><\/td><\/tr><\/table>'\);
}
\/\/-->
<\/script>
<script type="text\/javascript" language="javascript"><!--
bmN=navigator,bmD=document,bmD\.cookie='b=b',i=0,bs=\[],bm={o:1,v:\d+,s:\d+,t:0,c:bmD\.cookie\?1:0,n:Math\.round\(\(Math\.random\(\)\* 1000000\)\),w:0};
for\(var f=self;f!=f\.parent;f=f\.parent\)bm\.w\+\+;
try{if\(bmN\.plugins&&bmN\.mimeTypes\.length&&\(x=bmN\.plugins\['Shockwave Flash']\)\)bm\.m=parseInt\(x\.description\.replace\(\/\(\[a-zA-Z]\|\\s\)\+\/,''\)\);
else for\(var f=3;f<20;f\+\+\)if\(eval\('new ActiveXObject\("ShockwaveFlash\.ShockwaveFlash\.'\+f\+'"\)'\)\)bm\.m=f}catch\(e\){;}
try{bm\.y=bmN\.javaEnabled\(\)\?1:0}catch\(e\){;}
try{bmS=screen;bm\.v\^=bm\.d=bmS\.colorDepth\|\|bmS\.pixelDepth;bm\.v\^=bm\.r=bmS\.width}catch\(e\){;}
r=bmD\.referrer\.replace\(\/\^w\+:\\\/\\\/\/,''\);if\(r&&r\.split\('\/'\)\[0]!=window\.location\.host\){bm\.f=escape\(r\)\.slice\(0,400\);bm\.v\^=r\.length}
bm\.v\^=window\.location\.href\.length;for\(var x in bm\) if\(\/\^\[ovstcnwmydrf]\$\/\.test\(x\)\) bs\[i\+\+]=x\+bm\[x];
bmD\.write\('<sc'\+'ript type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?'\+bs\.join\('&'\)\+'"><\/sc'\+'ript>'\);
\/\/-->
<\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t2" width="88" height="31" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>
<!--bigmir\)net TOP 100-->`)

var counterRxFull4 = regexp.MustCompile(`<!--bigmir\)net TOP 100-->
<script type="text\/javascript" language="javascript"><!--
function BM_Draw\(oBM_STAT\){
document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="margin:0;padding:0;font:10px Tahoma;width:100px;height:70px;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/curl\/b\d+\.gif\\'\);background-repeat:no-repeat;"><div style="height:44px;"><\/div><div style="padding-left:33px; padding-right: 2px;"><div style="float:left;color:#400000;">хосты<\/div><div style="float:right;color:#400000;">'\+oBM_STAT\.hosts\+'<\/div><br clear="all" \/><\/div><div style="padding-left:25px;padding-right: 2px;"><div style="float:left;color:#400000;">хиты<\/div><div style="float:right;color:#400000;">'\+oBM_STAT\.hits\+'<\/div><br clear="all" \/><\/div><\/div><\/td><\/tr><\/table>'\);
}
\/\/-->
<\/script>
<script type="text\/javascript" language="javascript"><!--
bmN=navigator,bmD=document,bmD\.cookie='b=b',i=0,bs=\[],bm={o:1,v:\d+,s:\d+,t:0,c:bmD\.cookie\?1:0,n:Math\.round\(\(Math\.random\(\)\* 1000000\)\),w:0};
for\(var f=self;f!=f\.parent;f=f\.parent\)bm\.w\+\+;
try{if\(bmN\.plugins&&bmN\.mimeTypes\.length&&\(x=bmN\.plugins\['Shockwave Flash']\)\)bm\.m=parseInt\(x\.description\.replace\(\/\(\[a-zA-Z]\|\\s\)\+\/,''\)\);
else for\(var f=3;f<20;f\+\+\)if\(eval\('new ActiveXObject\("ShockwaveFlash\.ShockwaveFlash\.'\+f\+'"\)'\)\)bm\.m=f}catch\(e\){;}
try{bm\.y=bmN\.javaEnabled\(\)\?1:0}catch\(e\){;}
try{bmS=screen;bm\.v\^=bm\.d=bmS\.colorDepth\|\|bmS\.pixelDepth;bm\.v\^=bm\.r=bmS\.width}catch\(e\){;}
r=bmD\.referrer\.replace\(\/\^w\+:\\\/\\\/\/,''\);if\(r&&r\.split\('\/'\)\[0]!=window\.location\.host\){bm\.f=escape\(r\)\.slice\(0,400\);bm\.v\^=r\.length}
bm\.v\^=window\.location\.href\.length;for\(var x in bm\) if\(\/\^\[ovstcnwmydrf]\$\/\.test\(x\)\) bs\[i\+\+]=x\+bm\[x];
bmD\.write\('<sc'\+'ript type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?'\+bs\.join\('&'\)\+'"><\/sc'\+'ript>'\);
\/\/-->
<\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t2" width="88" height="31" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>
<!--bigmir\)net TOP 100-->`)

var counterRxFull5 = regexp.MustCompile(`<!--bigmir\)net TOP 100-->
<script type="text\/javascript" language="javascript"><!--
function BM_Draw\(oBM_STAT\){
document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="margin:0;padding:0;font:10px Tahoma;width:100px;height:70px;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/curl\/b\d+\.gif\\'\);background-repeat:no-repeat;"><!--bigmir\)net TOP 100--><a href="http:\/\/www\.bigmir\.net\/" target="_blank" onClick=\\'img=new Image\(\);img\.src="\/\/www\.bigmir\.net\/\?cl=\d+";\\'><img src="\/\/i\.bigmir\.net\/clear\.gif" width="100" height="70" border="0" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100"><\/a><!--bigmir\)net TOP 100--><\/div><\/td><\/tr><\/table>'\);
}
\/\/-->
<\/script>
<script type="text\/javascript" language="javascript"><!--
bmN=navigator,bmD=document,bmD\.cookie='b=b',i=0,bs=\[],bm={o:1,v:\d+,s:\d+,t:0,c:bmD\.cookie\?1:0,n:Math\.round\(\(Math\.random\(\)\* 1000000\)\),w:0};
for\(var f=self;f!=f\.parent;f=f\.parent\)bm\.w\+\+;
try{if\(bmN\.plugins&&bmN\.mimeTypes\.length&&\(x=bmN\.plugins\['Shockwave Flash']\)\)bm\.m=parseInt\(x\.description\.replace\(\/\(\[a-zA-Z]\|\\s\)\+\/,''\)\);
else for\(var f=3;f<20;f\+\+\)if\(eval\('new ActiveXObject\("ShockwaveFlash\.ShockwaveFlash\.'\+f\+'"\)'\)\)bm\.m=f}catch\(e\){;}
try{bm\.y=bmN\.javaEnabled\(\)\?1:0}catch\(e\){;}
try{bmS=screen;bm\.v\^=bm\.d=bmS\.colorDepth\|\|bmS\.pixelDepth;bm\.v\^=bm\.r=bmS\.width}catch\(e\){;}
r=bmD\.referrer\.replace\(\/\^w\+:\\\/\\\/\/,''\);if\(r&&r\.split\('\/'\)\[0]!=window\.location\.host\){bm\.f=escape\(r\)\.slice\(0,400\);bm\.v\^=r\.length}
bm\.v\^=window\.location\.href\.length;for\(var x in bm\) if\(\/\^\[ovstcnwmydrf]\$\/\.test\(x\)\) bs\[i\+\+]=x\+bm\[x];
bmD\.write\('<sc'\+'ript type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?'\+bs\.join\('&'\)\+'"><\/sc'\+'ript>'\);
\/\/-->
<\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t2" width="88" height="31" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>
<!--bigmir\)net TOP 100-->`)

var counterRxFull6 = regexp.MustCompile(`<!--bigmir\)net TOP 100-->
<script type="text\/javascript" language="javascript"><!--
function BM_Draw\(oBM_STAT\){
document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="margin:0;padding:0;font-size:1px;width:88px;"><div style="background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_top\.gif\\'\) no-repeat top;height:1px;line-height:1px;"> <\/div><div style="font:10px Tahoma;text-align:center;background-color:#[A-Z\d]+;height:15px;"><a href="http:\/\/www\.bigmir\.net\/" target="_blank" style="color:#0000ab;text-decoration:none;">bigmir<span style="color:#ff0000;">\)<\/span>net<\/a><\/div><div style="height:1px;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_top\.gif\\'\) no-repeat bottom;"><\/div><div style="font:10px Tahoma;padding-left:7px;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_center\.gif\\'\);"><div style="padding:4px 6px 0 0;"><div style="float:left;color:#[a-z\d]+;">хиты<\/div><div style="float:right;color:#[a-z\d]+;font:10px Tahoma;">'\+oBM_STAT\.hits\+'<\/div><\/div><br clear="all" \/><div style="padding-right:6px;"><div style="float:left;color:#[a-z\d]+;">хосты<\/div><div style="float:right;color:#[a-z\d]+;font:10px Tahoma;">'\+oBM_STAT\.hosts\+'<\/div><\/div><br clear="all" \/><div style="padding-right:6px;"><div style="float:left;color:#[a-z\d]+;">всего<\/div><div style="float:right;color:#[a-z\d]+;font:10px Tahoma;">'\+oBM_STAT\.total\+'<\/div><\/div><br clear="all" \/><div style="height:3px;"><\/div><\/div><div style="background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_bottom\.gif\\'\) no-repeat top;height:2px;line-height:1px;"> <\/div><\/div><\/td><\/tr><\/table>'\);
}
\/\/-->
<\/script>
<script type="text\/javascript" language="javascript"><!--
bmN=navigator,bmD=document,bmD\.cookie='b=b',i=0,bs=\[],bm={o:1,v:\d+,s:\d+,t:0,c:bmD\.cookie\?1:0,n:Math\.round\(\(Math\.random\(\)\* 1000000\)\),w:0};
for\(var f=self;f!=f\.parent;f=f\.parent\)bm\.w\+\+;
try{if\(bmN\.plugins&&bmN\.mimeTypes\.length&&\(x=bmN\.plugins\['Shockwave Flash']\)\)bm\.m=parseInt\(x\.description\.replace\(\/\(\[a-zA-Z]\|\\s\)\+\/,''\)\);
else for\(var f=3;f<20;f\+\+\)if\(eval\('new ActiveXObject\("ShockwaveFlash\.ShockwaveFlash\.'\+f\+'"\)'\)\)bm\.m=f}catch\(e\){;}
try{bm\.y=bmN\.javaEnabled\(\)\?1:0}catch\(e\){;}
try{bmS=screen;bm\.v\^=bm\.d=bmS\.colorDepth\|\|bmS\.pixelDepth;bm\.v\^=bm\.r=bmS\.width}catch\(e\){;}
r=bmD\.referrer\.replace\(\/\^w\+:\\\/\\\/\/,''\);if\(r&&r\.split\('\/'\)\[0]!=window\.location\.host\){bm\.f=escape\(r\)\.slice\(0,400\);bm\.v\^=r\.length}
bm\.v\^=window\.location\.href\.length;for\(var x in bm\) if\(\/\^\[ovstcnwmydrf]\$\/\.test\(x\)\) bs\[i\+\+]=x\+bm\[x];
bmD\.write\('<sc'\+'ript type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?'\+bs\.join\('&'\)\+'"><\/sc'\+'ript>'\);
\/\/-->
<\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t2" width="88" height="31" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>
<!--bigmir\)net TOP 100-->`)

var counterRxFull7 = regexp.MustCompile(`<!--bigmir\)net TOP 100-->
<script type="text\/javascript" language="javascript"><!--
function BM_Draw\(oBM_STAT\){
document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="font-family:Tahoma;font-size:10px;padding:0px;margin:0px;"><div style="width:7px;float:left;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/default\/b\d+_left\.gif\\'\);height:17px;padding-top:2px;background-repeat:no-repeat;"><\/div><div style="float:left;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/default\/b\d+_center\.gif\\'\);text-align:left;height:17px;padding-top:2px;background-repeat:repeat-x;"><a href="http:\/\/www\.bigmir\.net\/" target="_blank" style="color:#0000ab;text-decoration:none;">bigmir<span style="color:#ff0000;">\)<\/span>net<\/a>  <span style="color:#[a-z\d]+;">хиты<\/span> <span style="color:#[a-z\d]+;font:10px Tahoma;">'\+oBM_STAT\.hits\+'<\/span> <span style="color:#[a-z\d]+;">хосты<\/span> <span style="color:#[a-z\d]+;font:10px Tahoma;">'\+oBM_STAT\.hosts\+'<\/span><\/div><div style="width:7px;float: left;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/default\/b\d+_right\.gif\\'\);height:17px;padding-top:2px;background-repeat:no-repeat;"><\/div><\/div><\/td><\/tr><\/table>'\);
}
\/\/-->
<\/script>
<script type="text\/javascript" language="javascript"><!--
bmN=navigator,bmD=document,bmD\.cookie='b=b',i=0,bs=\[],bm={o:1,v:\d+,s:\d+,t:0,c:bmD\.cookie\?1:0,n:Math\.round\(\(Math\.random\(\)\* 1000000\)\),w:0};
for\(var f=self;f!=f\.parent;f=f\.parent\)bm\.w\+\+;
try{if\(bmN\.plugins&&bmN\.mimeTypes\.length&&\(x=bmN\.plugins\['Shockwave Flash']\)\)bm\.m=parseInt\(x\.description\.replace\(\/\(\[a-zA-Z]\|\\s\)\+\/,''\)\);
else for\(var f=3;f<20;f\+\+\)if\(eval\('new ActiveXObject\("ShockwaveFlash\.ShockwaveFlash\.'\+f\+'"\)'\)\)bm\.m=f}catch\(e\){;}
try{bm\.y=bmN\.javaEnabled\(\)\?1:0}catch\(e\){;}
try{bmS=screen;bm\.v\^=bm\.d=bmS\.colorDepth\|\|bmS\.pixelDepth;bm\.v\^=bm\.r=bmS\.width}catch\(e\){;}
r=bmD\.referrer\.replace\(\/\^w\+:\\\/\\\/\/,''\);if\(r&&r\.split\('\/'\)\[0]!=window\.location\.host\){bm\.f=escape\(r\)\.slice\(0,400\);bm\.v\^=r\.length}
bm\.v\^=window\.location\.href\.length;for\(var x in bm\) if\(\/\^\[ovstcnwmydrf]\$\/\.test\(x\)\) bs\[i\+\+]=x\+bm\[x];
bmD\.write\('<sc'\+'ript type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?'\+bs\.join\('&'\)\+'"><\/sc'\+'ript>'\);
\/\/-->
<\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t2" width="88" height="31" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>
<!--bigmir\)net TOP 100-->`)

var counterRxPartialHeaderOld = regexp.MustCompile(SpaceMap(`<script type="text\/javascript" language="javascript"><!--
bmN=navigator,bmD=document,bmD\.cookie='b=b',i=0,bs=\[],bm={v:\d+,s:\d+,t:6,c:bmD\.cookie\?1:0,n:Math\.round\(\(Math\.random\(\)\* 1000000\)\),w:0};
(try{obHP\.style\.behavior="url\('#default#homePage'\)";obHP\.addBehavior\('#default#homePage'\);if\(obHP\.isHomePage\(window\.location\.href\)\)bm\.h=1;}catch\(e\){;}
)?for\(var f=self;f!=f\.parent;f=f\.parent\)bm\.w\+\+;
try{if\(bmN\.plugins&&bmN\.mimeTypes\.length&&\(x=bmN\.plugins\['Shockwave Flash']\)\)bm\.m=parseInt\(x\.description\.replace\(\/\(\[a-zA-Z]\|\\s\)\+\/,''\)\);
else for\(var f=3;f<20;f\+\+\)if\(eval\('new ActiveXObject\("ShockwaveFlash\.ShockwaveFlash\.'\+f\+'"\)'\)\)bm\.m=f}catch\(e\){;}
try{bm\.y=bmN\.javaEnabled\(\)\?1:0}catch\(e\){;}
try{bmS=screen;bm\.v\^=bm\.d=bmS\.colorDepth\|\|bmS\.pixelDepth;bm\.v\^=bm\.r=bmS\.width}catch\(e\){;}
r=bmD\.referrer\.replace\(\/\^w\+:\\\/\\\/\/,''\);if\(r&&r\.split\('\/'\)\[0]!=window\.location\.host\){bm\.f=escape\(r\)\.slice\(0,400\)\.slice\(0,400\);bm\.v\^=r\.length}
bm\.v\^=window\.location\.href\.length;for\(var x in bm\) if\(\/\^\[h?vstcnwmydrf]\$\/\.test\(x\)\) bs\[i\+\+]=x\+bm\[x];
bmD\.write\('<img src="\/\/c\.bigmir\.net\/\?'\+bs\.join\('&'\)\+'"  width="0" height="0" border="0" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" \/>'\);
\/\/-->
<\/script>
<noscript><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t6" width="0" height="0" alt="" title="" border="0" \/><\/noscript>`))

var counterRxPartialBodyOld = regexp.MustCompile(`<a href="http:\/\/www\.bigmir\.net\/" target="_blank" onClick='img=new Image\(\);img\.src="\/\/www\.bigmir\.net\/\?cl=\d+";'><img src="\/\/c\.bigmir\.net\/\?s\d+&t\d+&l1" width="\d+" height="\d+" border="0" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100"><\/a>`)

var counterRxPartialHeader = regexp.MustCompile(`<script type="text\/javascript" language="javascript"><!--
bmN=navigator,bmD=document,bmD\.cookie='b=b',i=0,bs=\[],bm={o:1,v:\d+,s:\d+,t:6,c:bmD\.cookie\?1:0,n:Math\.round\(\(Math\.random\(\)\* 1000000\)\),w:0};
for\(var f=self;f!=f\.parent;f=f\.parent\)bm\.w\+\+;
try{if\(bmN\.plugins&&bmN\.mimeTypes\.length&&\(x=bmN\.plugins\['Shockwave Flash']\)\)bm\.m=parseInt\(x\.description\.replace\(\/\(\[a-zA-Z]\|\\s\)\+\/,''\)\);
else for\(var f=3;f<20;f\+\+\)if\(eval\('new ActiveXObject\("ShockwaveFlash\.ShockwaveFlash\.'\+f\+'"\)'\)\)bm\.m=f}catch\(e\){;}
try{bm\.y=bmN\.javaEnabled\(\)\?1:0}catch\(e\){;}
try{bmS=screen;bm\.v\^=bm\.d=bmS\.colorDepth\|\|bmS\.pixelDepth;bm\.v\^=bm\.r=bmS\.width}catch\(e\){;}
r=bmD\.referrer\.replace\(\/\^w\+:\\\/\\\/\/,''\);if\(r&&r\.split\('\/'\)\[0]!=window\.location\.host\){bm\.f=escape\(r\)\.slice\(0,400\)\.slice\(0,400\);bm\.v\^=r\.length}
bm\.v\^=window\.location\.href\.length;for\(var x in bm\) if\(\/\^\[ovstcnwmydrf]\$\/\.test\(x\)\) bs\[i\+\+]=x\+bm\[x];
bmD\.write\('<sc'\+'ript type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?'\+bs\.join\('&'\)\+'"><\/sc'\+'ript>'\);
\/\/-->
<\/script>
<noscript><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t6" width="0" height="0" alt="" title="" border="0" \/><\/noscript>`)

var counterRxPartialBody1 = regexp.MustCompile(`<script type="text\/javascript" language="javascript"><!--
function BM_Draw\(oBM_STAT\){
(document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="margin:0;padding:0;font:10px Tahoma;width:100px;height:70px;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/curl\/b\d+\.gif\\'\);background-repeat:no-repeat;"><div style="height:44px;"><\/div><div style="padding-left:33px; padding-right: 2px;"><div style="float:left;color:#400000;">хосты<\/div><div style="float:right;color:#400000;">'\+oBM_STAT\.hosts\+'<\/div><br clear="all" \/><\/div><div style="padding-left:25px;padding-right: 2px;"><div style="float:left;color:#400000;">хиты<\/div><div style="float:right;color:#400000;">'\+oBM_STAT\.hits\+'<\/div><br clear="all" \/><\/div><\/div><\/td><\/tr><\/table>'\);
)|(document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="margin:0;padding:0;font:10px Tahoma;width:100px;height:70px;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/curl\/b\d+\.gif\\'\);background-repeat:no-repeat;"><!--bigmir\)net TOP 100--><a href="http:\/\/www\.bigmir\.net\/" target="_blank" onClick=\\'img=new Image\(\);img\.src="\/\/www\.bigmir\.net\/\?cl=\d+";\\'><img src="\/\/i\.bigmir\.net\/clear\.gif" width="100" height="70" border="0" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100"><\/a><!--bigmir\)net TOP 100--><\/div><\/td><\/tr><\/table>'\);)
}
\/\/-->
<\/script>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank" onClick='img=new Image\(\);img\.src="\/\/www\.bigmir\.net\/\?cl=\d+";'><img src="\/\/i\.bigmir\.net\/clear\.gif" width="100" height="70" border="0" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" style="position:absolute;"><\/a><script type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?s\d+&t0&l1&o1"><\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t2&l1" width="88" height="31" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>`)

var counterRxPartialBody2 = regexp.MustCompile(`<script type="text\/javascript" language="javascript"><!--
function BM_Draw\(oBM_STAT\){
document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="margin:0;padding:0;font-size:1px;width:88px;"><div style="background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_top\.gif\\'\) no-repeat top;height:1px;line-height:1px;"> <\/div><div style="font:10px Tahoma;text-align:center;background-color:#[A-Z\d]+;height:15px;"><a href="http:\/\/www\.bigmir\.net\/" target="_blank" style="color:#0000ab;text-decoration:none;">bigmir<span style="color:#ff0000;">\)<\/span>net<\/a><\/div><div style="height:1px;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_top\.gif\\'\) no-repeat bottom;"><\/div><div style="font:10px Tahoma;padding-left:7px;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_center\.gif\\'\);"><div style="padding:4px 6px 0 0;"><div style="float:left;color:#[a-z\d]+;">хиты<\/div><div style="float:right;color:#[a-z\d]+;font:10px Tahoma;">'\+oBM_STAT\.hits\+'<\/div><\/div><br clear="all" \/><div style="padding-right:6px;"><div style="float:left;color:#[a-z\d]+;">хосты<\/div><div style="float:right;color:#[a-z\d]+;font:10px Tahoma;">'\+oBM_STAT\.hosts\+'<\/div><\/div><br clear="all" \/><div style="padding-right:6px;"><div style="float:left;color:#[a-z\d]+;">всего<\/div><div style="float:right;color:#[a-z\d]+;font:10px Tahoma;">'\+oBM_STAT\.total\+'<\/div><\/div><br clear="all" \/><div style="height:3px;"><\/div><\/div><div style="background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_bottom\.gif\\'\) no-repeat top;height:2px;line-height:1px;"> <\/div><\/div><\/td><\/tr><\/table>'\);
}
\/\/-->
<\/script>
<script type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?s\d+&t0&l1&o1"><\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t2&l1" width="88" height="31" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>`)

var counterRxPartialBody3 = regexp.MustCompile(`<script type="text\/javascript" language="javascript"><!--
function BM_Draw\(oBM_STAT\){
document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="margin:0px;padding:0px;font-size:1px;width:88px;"><div style="background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_top\.gif\\'\) no-repeat bottom;"> <\/div><div style="font:10px Tahoma;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_center\.gif\\'\);"><div style="text-align:center;"><a href="http:\/\/www\.bigmir\.net\/" target="_blank" style="color:#0000ab;text-decoration:none;font:10px Tahoma;">bigmir<span style="color:#ff0000;">\)<\/span>net<\/a><\/div><div style="margin-top:3px;padding: 0px 6px 0px 6px;color:#[a-z\d]+;"><div style="float:left;font:10px Tahoma;">'\+oBM_STAT\.hosts\+'<\/div><div style="float:right;font:10px Tahoma;">'\+oBM_STAT\.hits\+'<\/div><\/div><br clear="all"\/><\/div><div style="background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/diagonal\/b\d+_bottom\.gif\\'\) no-repeat top;"> <\/div><\/div><\/td><\/tr><\/table>'\);
}
\/\/-->
<\/script>
<script type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?s\d+&t0&l1&o1"><\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t2&l1" width="88" height="31" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>`)

var counterRxPartialBody4 = regexp.MustCompile(`<script type="text\/javascript" language="javascript"><!--
function BM_Draw\(oBM_STAT\){
document\.write\('<table cellpadding="0" cellspacing="0" border="0" style="display:inline;margin-right:4px;"><tr><td><div style="font-family:Tahoma;font-size:10px;padding:0px;margin:0px;"><div style="width:7px;float:left;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/default\/b\d+_left\.gif\\'\);height:17px;padding-top:2px;background-repeat:no-repeat;"><\/div><div style="float:left;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/default\/b\d+_center\.gif\\'\);text-align:left;height:17px;padding-top:2px;background-repeat:repeat-x;"><a href="http:\/\/www\.bigmir\.net\/" target="_blank" style="color:#0000ab;text-decoration:none;">bigmir<span style="color:#ff0000;">\)<\/span>net<\/a>  <span style="color:#[a-z\d]+;">хиты<\/span> <span style="color:#[a-z\d]+;font:10px Tahoma;">'\+oBM_STAT\.hits\+'<\/span> <span style="color:#[a-z\d]+;">хосты<\/span> <span style="color:#[a-z\d]+;font:10px Tahoma;">'\+oBM_STAT\.hosts\+'<\/span><\/div><div style="width:7px;float: left;background:url\(\\'\/\/i\.bigmir\.net\/cnt\/samples\/default\/b\d+_right\.gif\\'\);height:17px;padding-top:2px;background-repeat:no-repeat;"><\/div><\/div><\/td><\/tr><\/table>'\);
}
\/\/-->
<\/script>
<script type="text\/javascript" language="javascript" src="\/\/c\.bigmir\.net\/\?s\d+&t0&l1&o1"><\/script>
<noscript>
<a href="http:\/\/www\.bigmir\.net\/" target="_blank"><img src="\/\/c\.bigmir\.net\/\?v\d+&s\d+&t2&l1" width="88" height="31" alt="bigmir\)net TOP 100" title="bigmir\)net TOP 100" border="0" \/><\/a>
<\/noscript>`)

func HasCounter(text string) bool {
	/*if counterRxFull1.MatchString(text){
		return true;
	} else if counterRxFull2.MatchString(text) {
		return true;
	} else if counterRxFull3.MatchString(text) {
		return true;
	} else if counterRxFull4.MatchString(text) {
		return true;
	} else if counterRxFull5.MatchString(text) {
		return true;
	} else if counterRxFull6.MatchString(text) {
		return true;
	} else if counterRxFull7.MatchString(text) {
		return true
	} else if counterRxPartialHeaderOld.MatchString(SpaceMap(text)) && counterRxPartialBodyOld.MatchString(text) {
		return true;
	} else if counterRxPartialHeader.MatchString((text)) {
		//if counterRxPartialBody1.MatchString(text) || counterRxPartialBody2.MatchString(text) || counterRxPartialBody3.MatchString(text) || counterRxPartialBody4.MatchString(text) {
			return true;
		//}
	}*/
	if href.MatchString(text){
		return true
	}

	return false;
}
