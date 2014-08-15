---
title: 'Using &#8216;background-position&#8217; and sprite sheets to stop icon hover flicker'
author: Jacob Tomlinson
layout: post
permalink: /2012/09/18/using-background-position-and-sprite-sheets-to-stop-icon-hover-flicker/
original-css-code:
  - |
    .facebook-icon {
    	height: 32px;
    	width: 32px;
    	background-image: url(images/icons/facebook_dark.png);
    }
    
    .facebook-icon:hover {
    	background-image: url(images/icons/facebook_active.png);
    }
new-css-code:
  - |
    .social-icon {
    	height: 32px;
    	width: 32px;
    	background-image: url(images/icons/social_icons.png);
    }
    
    .facebook-icon {
    	background-position: -32px 0px;
    }
    
    .facebook-icon:hover {
    	background-position: -32px -32px;
    }
original-div-code:
  - '<div class="facebook-icon"></div>'
new-div-code:
  - '<div class="social-icon facebook-icon"></div>'
has_been_twittered:
  - yes
categories:
  - Web Development
tags:
  - CSS
  - Guide
  - HTML
---
While updating the theme on this blog I added some links to my social websites. I made these links in the form of images which were black and white and become coloured when hovered over. To create these icons I created empty divs which would then be styled in CSS.

<pre class="lang:default decode:true">&lt;div class="facebook-icon"&gt;&lt;/div&gt;</pre>

The CSS set the div size and the background image and then changed the background image on hover. Something along the lines of

<pre class="lang:default decode:true">.facebook-icon {
	height: 32px;
	width: 32px;
	background-image: url(images/icons/facebook_dark.png);
}

.facebook-icon:hover {
	background-image: url(images/icons/facebook_active.png);
}</pre>

Now this was fine and worked well but it had a problem. When the user hovered over the button for the first time it would flicker because the &#8216;active&#8217; image hadn&#8217;t been loaded yet and it had to wait for the request to return the image. This made the buttons look sloppy and so I had to rethink my approach.

To fix it I used a very basic web design technique which involves putting all the icon images into one file called a sprite sheet and then shifting the background around using the &#8216;background-position&#8217; property in CSS. This way the page only requests one file which is loaded at the beginning to show the darker images and therefore the colour images are loaded but just not shown.

The best way to think about it is to imagine the div as a little window that you&#8217;re looking at the sprite sheet through. You can only see a small part the size of the window and if you move the sprite sheet around the image in the window changes.

<a style="font-style: normal; line-height: 24px; text-decoration: underline;" href="http://www.jacobtomlinson.co.uk/2012/09/18/using-background-position-and-sprite-sheets-to-stop-icon-hover-flicker/social_icons_demo_1/" rel="attachment wp-att-204"><img class="size-full wp-image-204 alignright" style="border-style: initial; border-color: initial; background-image: initial; background-attachment: initial; background-origin: initial; background-clip: initial; background-color: #eeeeee; margin-top: 0.4em;" title="Social icons position example 1" alt="" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2012/09/social_icons_demo_1.png" width="128" height="64" /></a>

So in this example think of the red square as the div which is 32px by 32px. To position the grey facebook icon inside the square the position has to be set to -32px 0px which shifts the background image 32 pixels to the left.

<a href="http://www.jacobtomlinson.co.uk/2012/09/18/using-background-position-and-sprite-sheets-to-stop-icon-hover-flicker/social_icons_demo_2/" rel="attachment wp-att-205"><img class="size-full wp-image-205 alignright" title="Social icons position demo 2" alt="" src="http://www.jacobtomlinson.co.uk/wp-content/uploads/2012/09/social_icons_demo_2.png" width="128" height="64" /></a>

Then when the image is hovered over the x position stays at -32px for the facebook icon but the y axis changes to -32px to shift the sprite sheet up 32 pixels and places the colour facebook icon inside the div.

So the CSS for this would be

<pre class="lang:default decode:true">.social-icon {
	height: 32px;
	width: 32px;
	background-image: url(images/icons/social_icons.png);
}

.facebook-icon {
	background-position: -32px 0px;
}

.facebook-icon:hover {
	background-position: -32px -32px;
}</pre>

Now there are a few changes in this code. One you will notice is that there is an extra class. Each icon now has both the default &#8216;social-icon&#8217; class which gives the size and background image and then it&#8217;s own class to set the position and hover class to set the hover position. To include the default class the div would be changed to

<pre class="lang:default decode:true  crayon-selected">&lt;div class="social-icon facebook-icon"&gt;&lt;/div&gt;</pre>

Then similar classes are created for &#8216;twitter-icon&#8217;, &#8216;linkedin-icon&#8217; and &#8216;googleplus-icon&#8217;. This means that not only does it save time in creating new icons as the height and width don&#8217;t have to be specified each time and solves the icons flickering but also saves many server requests per page.