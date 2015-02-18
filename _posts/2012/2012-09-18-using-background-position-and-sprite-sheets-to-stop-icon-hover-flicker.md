---
title: 'Using background-position and sprite sheets to stop icon hover flicker'
author: Jacob Tomlinson
layout: post
permalink: /2012/09/18/using-background-position-and-sprite-sheets-to-stop-icon-hover-flicker/
category: Web Development
thumbnail: internet
tags:
  - CSS
  - Guide
  - HTML
---

While updating the theme on this blog I added some links to my social websites. I made these links in the form of images which were black and white and become coloured when hovered over. To create these icons I created empty divs which would then be styled in CSS.

```html
<div class="facebook-icon"></div>
```

The CSS set the div size and the background image and then changed the background image on hover. Something along the lines of

```css
.facebook-icon {
	height: 32px;
	width: 32px;
	background-image: url(images/icons/facebook_dark.png);
}

.facebook-icon:hover {
	background-image: url(images/icons/facebook_active.png);
}
```

This was fine and worked well but it had a problem. When the user hovered over the button for the first time it would flicker because the 'active' image hadn't been loaded yet and it had to wait for the request to return the image. This made the buttons look sloppy and so I had to rethink my approach.

To fix it I used a very basic web design technique which involves putting all the icon images into one file called a sprite sheet and then shifting the background around using the 'background-position' property in CSS. This way the page only requests one file which is loaded at the beginning to show the darker images and therefore the colour images are loaded but just not shown.

The best way to think about it is to imagine the div as a little window that you're looking at the sprite sheet through. You can only see a small part the size of the window and if you move the sprite sheet around the image in the window changes.

![Icon Sprite Sheet Grey](http://i.imgur.com/PZq9a1m.png)

So in this example think of the red square as the div which is 32px by 32px. To position the grey facebook icon inside the square the position has to be set to -32px 0px which shifts the background image 32 pixels to the left.

![Icon Sprite Sheet Colour](http://i.imgur.com/oI04LAo.png)

Then when the image is hovered over the x position stays at -32px for the facebook icon but the y axis changes to -32px to shift the sprite sheet up 32 pixels and places the colour facebook icon inside the div.

So the CSS for this would be

```css
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
```

There are a few changes in this code. One you will notice is that there is an extra class. Each icon now has both the default 'social-icon' class which gives the size and background image and then it's own class to set the position and hover class to set the hover position. To include the default class the div would be changed to

```html
<div class="social-icon facebook-icon"></div></pre>
```

Then similar classes are created for `twitter-icon`, `linkedin-icon` and `googleplus-icon`. This means that not only does it save time in creating new icons as the height and width don't have to be specified each time and solves the icons flickering but also saves many server requests per page.
