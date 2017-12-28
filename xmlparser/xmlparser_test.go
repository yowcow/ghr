package xmlparser

import (
	"testing"
)

var AtomXML = `
<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" xml:lang="en-US">
  <id>tag:github.com,2008:https://github.com/vim/vim/releases</id>
  <link type="text/html" rel="alternate" href="https://github.com/vim/vim/releases"/>
  <link type="application/atom+xml" rel="self" href="https://github.com/vim/vim/releases.atom"/>
  <title>Release notes from vim</title>
  <updated>2017-08-21T20:06:02Z</updated>
  <entry>
    <id>tag:github.com,2008:Repository/40997482/v8.0.0983</id>
    <updated>2017-08-21T20:06:02Z</updated>
    <link rel="alternate" type="text/html" href="/vim/vim/releases/tag/v8.0.0983"/>
    <title>v8.0.0983: patch 8.0.0983: unnecessary check for NULL pointer</title>
    <content type="html">&lt;p&gt;Problem:    Unnecessary check for NULL pointer.&lt;br&gt;
Solution:   Remove the NULL check in dialog_changed(), it already happens in&lt;br&gt;
dialog_msg(). (Ken Takata)&lt;/p&gt;</content>
    <author>
      <name>brammool</name>
    </author>
    <media:thumbnail height="30" width="30" url="https://avatars0.githubusercontent.com/u/8530623?v=4&amp;s=60"/>
  </entry>
  <entry>
    <id>tag:github.com,2008:Repository/40997482/v8.0.0982</id>
    <updated>2017-08-21T20:01:27Z</updated>
    <link rel="alternate" type="text/html" href="/vim/vim/releases/tag/v8.0.0982"/>
    <title>v8.0.0982: patch 8.0.0982: cannot use a terminal when &#39;encoding&#39; is non-utf8 mul…</title>
    <content type="html">&lt;p&gt;…ti-byte&lt;/p&gt;
&lt;p&gt;Problem:    When &#39;encoding&#39; is set to a multi-byte encoding other than utf-8&lt;br&gt;
the characters from ther terminal are messed up.&lt;br&gt;
Solution:   Convert displayed text from utf-8 to &#39;encoding&#39; for MS-Windows.&lt;br&gt;
(Yasuhiro Matsumoto, close &lt;a href=&quot;https://github.com/vim/vim/pull/2000&quot; class=&quot;issue-link js-issue-link&quot; data-url=&quot;https://github.com/vim/vim/issues/2000&quot; data-id=&quot;251691887&quot; data-error-text=&quot;Failed to load issue title&quot; data-permission-text=&quot;Issu
e title is private&quot;&gt;#2000&lt;/a&gt;)&lt;/p&gt;</content>
    <author>
      <name>brammool</name>
    </author>
    <media:thumbnail height="30" width="30" url="https://avatars0.githubusercontent.com/u/8530623?v=4&amp;s=60"/>
  </entry>
  <entry>
    <id>tag:github.com,2008:Repository/40997482/v8.0.0981</id>
    <updated>2017-08-21T19:39:28Z</updated>
    <link rel="alternate" type="text/html" href="/vim/vim/releases/tag/v8.0.0981"/>
    <title>v8.0.0981: patch 8.0.0981: cursor in terminal window blinks by default</title>
    <content type="html">&lt;p&gt;Problem:    Cursor in terminal window blinks by default, while in a real xterm&lt;br&gt;
it does not blink, unless the -bc argument is used.&lt;br&gt;
Solution:   Do not use a blinking cursor by default.&lt;/p&gt;</content>
    <author>
      <name>brammool</name>
    </author>
    <media:thumbnail height="30" width="30" url="https://avatars0.githubusercontent.com/u/8530623?v=4&amp;s=60"/>
  </entry>
</feed>
`

var InvalidXML = `
hoge <fuga>
`

func Test_ParseAtom_returns_3_elements(t *testing.T) {
	data := []byte(AtomXML)
	entries, err := ParseAtom(data)

	if err != nil {
		t.Error("Expected nil but got", err)
	}
	if len := len(entries); len != 3 {
		t.Error("Expected 3 elements but got", len)
	}
}

func Test_ParseAtom_returns_error(t *testing.T) {
	data := []byte(InvalidXML)
	entries, err := ParseAtom(data)

	if err == nil {
		t.Error("Expected error but got", err)
	}
	if entries != nil {
		t.Error("Expected nil but got", entries)
	}
}
