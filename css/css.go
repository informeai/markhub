package css

var TemplCss = `
:root{
	--titles: #24292E;
	--subtitles: #6A737D;
	--link-colors: #0366D6;
	--pre-background: #F6F8FA;
	--code: #1b1f230d;
	--border-table: #e0e3e6;
	--header-table: #24292e;
	--body-table: #14292e;
}

*{
	margin: 0;
	padding: 1rem 0;
	position: relative;
	box-sizing: border-box;
	color: var(--titles)
}
body{
	padding: 2rem;
	font-family: -apple-system,BlinkMacSystemFont,Segoe UI,Helvetica,Arial,sans-serif,Apple Color Emoji,Segoe UI Emoji;
	line-height: 1.5;
}

h1,h2,h3,h4,h5{
	color: var(--titles);
}

h6{
	color: var(--subtitles);
}

h1:after,h2:after{
	content: "";
	position: absolute;
	width: 100%;
	height: 1px;
	background-color: rgba(0,0,0,0.1);
	bottom: 0;
	left: 0;
}

blockquote{
	color: var(--subtitles);
	padding: 0;
	margin: 0;
}
blockquote > p{
	padding: 0;
	margin-left: 1rem;
	color: var(--subtitles);
}
blockquote:after{
	content: "";
	position: absolute;
	top: 0;
	left: 0;
	height: 100%;
	width: 3px;
	background-color: var(--subtitles);
}
a{
	color: var(--link-colors);
	text-decoration: none;
}

code{
	background-color: var(--code);
	padding: 0.3rem;
	border-radius: 0.3rem;
}
pre{
	padding: 1rem;
	color: var(--titles);
	background-color: var(--pre-background);
}
pre > code{
	background-color: var(--pre-background);
}
ul, ol{
	margin: 0;
	padding: 0;
	margin: 0 0 1rem;
	padding-left: 2rem;
}
ul > li, ol > li{
	margin: 0;
	padding: 0;
}
img{
	width: 100%;
	height: auto;
	object-fit: cover;
}
table{
	margin: 0;
	padding: 0;
	overflow: auto;
	width: 100%
	max-width: 100%;
	border-collapse: collapse; 
	border-spacing: 0; 
	/* border: 1px solid; */
	/* border-color: var(--border-table); */
}


thead{
	margin: 0;
	padding: 0;
	color: var(--header-table);
	display: table-header-group;
	vertical-align: middle;
	border: 1px solid var(--border-table);
}
tbody{
	padding: 0.3rem;
	color: var(--body-table);
	
}
thead  tr{
	margin: 0;
	padding: 0;
	display: table-row;
	vertical-align: inherit;
	border: 1px solid var(--border-table);
}
thead tr th{
	border: 1px solid var(--border-table);
	padding: 0.5rem 1rem;
}

tbody tr td{
	border: 1px solid var(--border-table);
	padding: 0.5rem 1rem;
}

`
