---
title: Using the AddThis Share Buttons wordpress plugin in a custom theme
author: Jacob Tomlinson
layout: post
permalink: /2013/03/17/using-the-addthis-share-buttons-wordpress-plugin-in-a-custom-theme/
thumbnail: add-this
category: Web Development
tags:
  - addthis
  - theme
  - wordpress
---

There is an undocumented function for adding a custom AddThis widget to your Wordpress theme when using the <a title="Add This Share Buttons WordPress Plugin" href="http://wordpress.org/extend/plugins/addthis/" target="_blank">Add This Share Buttons plugin</a>, so I thought I would document it here.

Usage:

```php
<?php do_action( 'addthis_widget', $the_permalink, $the_title, $custom_style ); ?>
```

Parameters:

**$the_permalink**  
(*string*) (*optional*) The url you want to share with AddThis.  
Default:  get_permalink()

**$the_title**  
(*string*) (*optional*) The title of the page you&#8217;re sharing.  
Default:  get_the_title()

**$custom_style**  
(*string*)  (*optional*) The name of the preset style you want to use         or  
(*array*) (*optional*) An array of the settings you want to use for your own custom style  
Default:  'fb_tw_p1_sc'

**Preset Styles**

<table class="table table-condensed">
  <tr>
    <td>
      <strong>Style</strong>
    </td>

    <td>
      <strong>Description</strong>
    </td>
  </tr>

  <tr>
    <td>
      <em>fb_tw_p1_sc</em>
    </td>

    <td>
      Facebook, twitter, pinterest and social counter
    </td>
  </tr>

  <tr>
    <td>
      <em>large_toolbox</em>
    </td>

    <td>
      Large toolbox
    </td>
  </tr>

  <tr>
    <td>
      <em>small_toolbox</em>
    </td>

    <td>
      Small toolbox
    </td>
  </tr>

  <tr>
    <td>
      <em>plus_one_share_counter</em>
    </td>

    <td>
      Google plus counter
    </td>
  </tr>

  <tr>
    <td>
      <em>small_toolbox_with_share</em>
    </td>

    <td>
      Small toolbox with share button
    </td>
  </tr>

  <tr>
    <td>
      <em>fb_tw_sc</em>
    </td>

    <td>
      Facebook, twitter and social counter
    </td>
  </tr>

  <tr>
    <td>
      <em>simple_button</em>
    </td>

    <td>
      Simple AddThis sharing button
    </td>
  </tr>

  <tr>
    <td>
      <em>button</em>
    </td>

    <td>
      Default AddThis sharing button
    </td>
  </tr>

  <tr>
    <td>
      <em>share_counter</em>
    </td>

    <td>
      AddThis share counter
    </td>
  </tr>

  <tr>
    <td>
      <em>above</em>
    </td>

    <td>
      SPECIAL loads your setting from the admin control panel for the top widget
    </td>
  </tr>

  <tr>
    <td>
      <em>below</em>
    </td>

    <td>
      SPECIAL loads your setting from the admin control panel for the bottom widget
    </td>
  </tr>
</table>

**Custom Style Settings**

<table class="table table-condensed">
  <tr>
    <td>
      <strong>Setting</strong>
    </td>

    <td>
      <strong>Description</strong>
    </td>
  </tr>

  <tr>
    <td>
      <em>type</em>
    </td>

    <td>
      (<em>string</em>) Must be set to <em>&#8216;custom&#8217;</em>
    </td>
  </tr>

  <tr>
    <td>
      <em>size</em>
    </td>

    <td>
      (<em>string</em>) Icon size, 16 or 32. e.g <em>&#8217;16&#8242;</em>
    </td>
  </tr>

  <tr>
    <td>
      <em>services</em>
    </td>

    <td>
      (<em>string</em>) Comma separated services. Full list of services <a title="AddThis Services" href="http://www.addthis.com/services/list#.UUX6QlqAuYQ" target="_blank">here</a>. e.g <em>&#8216;facebook, twitter, google_plusone_share, email, pinterest_share&#8217;</em>
    </td>
  </tr>

  <tr>
    <td>
      <em>preferred</em>
    </td>

    <td>
      (<em>string</em>) Number of additional services chosen from the user&#8217;s most used. e.g <em>&#8217;3&#8242;</em>
    </td>
  </tr>

  <tr>
    <td>
      <em>more</em>
    </td>

    <td>
      (<em>boolean</em>) If you want to have the more button at the end. e.g <em>true</em>
    </td>
  </tr>

  <tr>
    <td>
      <em>counter</em>
    </td>

    <td>
      (<em>boolean/string</em>) If you want the counter at the end e.g  <em>true</em> or the counter style (assumes true) e.g  <em>&#8216;bubble_style&#8217;</em>
    </td>
  </tr>
</table>

**Example Usage**

```php
<?php do_action( 'addthis_widget', get_permalink(), get_the_title(), 'small_toolbox'); ?>
```

```php
<?php do_action( 'addthis_widget', get_permalink(), get_the_title(), array(
    'type' => 'custom',
    'size' => '16', // size of the icons.  Either 16 or 32
    'services' => 'facebook,twitter,digg', // the services you want to always appear
    'preferred' => '7', // the number of auto personalized services
    'more' => true, // if you want to have a more button at the end
    'counter' => 'bubble_style' // if you want a counter and the style of it
    ));
?>
```

### How I found this

So I recently decided to give my website a bit of a facelift and as part of that I decided to add some better quality sharing buttons to my blog posts and portfolio projects. I settled on using the <a title="Add This Share Buttons wordpress plugin" href="http://wordpress.org/extend/plugins/addthis/" target="_blank">AddThis Share Buttons plugin</a> for WordPress. However when using the plugin I had a bit of trouble placing the AddThis widget exactly where I wanted it.

The main issue I came across was that when using the settings page for the plugin it only gave me the option to place it above and/or below my post and gave me a selection of a few different styles. When looking at the screenshots I saw buttons for custom and advanced settings and so I assumed I could easily place the widget where I wanted it and would be able to style it as simply as you can with the website version of the <a title="Add This Share Buttons" href="https://www.addthis.com/get/sharing" target="_blank">AddThis Share Buttons</a>. But once I actually installed the plugin and had a look through the settings I saw that the custom option just allowed you to specify how the box looks but not it&#8217;s location and the advanced settings were more for the analytics and other options.

By default if you have a sharing widget at the top the plugin hooks to <a title="Wordpress function the_content()" href="https://developer.wordpress.org/reference/functions/the_content/" target="_blank">the content</a> and places the code before your content and if you select bottom it places the code after your content. I didn&#8217;t want either of those options as I wanted my top widget to be in with the header in a div that I could specify and control and I wanted the bottom widget to be after other elements like the tags which I had specified in my theme and I wanted them to have a sharing icon which was in keeping with the tag icon that I&#8217;d given that tags.

So I began looking for either a php function or shortcode that I could put into my theme. I was unable to find anything in the plugin documentation so naturally I took to google to find a solution. The only thing I managed to find was <a title="Wordpress forum for Addthis custom widget" href="http://wordpress.org/support/topic/is-there-a-shortcode-for-addthis" target="_blank">this forum page</a> where someone suggested using the function

```php
<?php do_action( 'addthis_widget' ); ?>
```

I found that putting that into your theme gives you a default widget in that position so this is a start.

The next thing I wanted to do was to style the widget how I wanted it. I like the small &#8216;bubble&#8217; style default widget the you can choose from the default menu but I also wanted one with more options on to go at the bottom. So again back to google, this time with the function name, and found <a title="add_this function custom parameter" href="http://wordpress.org/support/topic/do_action-addthis_widget-customize" target="_blank">this forum page</a> where it was suggested that you could give an array as the fourth parameter of the action which would take custom options like this

```php
<?php do_action( 'addthis_widget', get_permalink(), get_the_title(), array(
    'size' => '16', // size of the icons.  Either 16 or 32
    'services' => 'hyves,joliprint', // the services you want to always appear
    'preferred' => '8', // the number of auto personalized services
    'more' => true // if you want to have a more button at the end
    ));
?>
```

The function takes the url to share and the title of the page as the second and third parameters so you have to specify them as we are trying to override the fourth. I had a little play with the options given in the example on that page but couldn&#8217;t get the plugin to look quite how I wanted.

I found a second solution where someone suggested that instead of passing an array you can pass a string with the name of the default style you want to use. They went on to suggest that you could then go into the plugin code, find the array of styles and add your own. I wouldn&#8217;t recommend this though as your changes would get written over next time the plugin is updated.  *Remember if you&#8217;re creating a custom theme keep all your changes in the theme directory or risk being overwritten in updates.*

But if you do want to just specify custom placement of the widget and not a custom style then this is a valid option, here is an example of the code and a list of all the default styles.

```php
<?php do_action( 'addthis_widget', get_permalink(), get_the_title(), 'small_toolbox' );
```

See table above for details.

But in order to try and specify the exact layout I decided to take a look into the plugin itself and find out what the function was doing. Now the first thing the function does is to check if you&#8217;ve passed it a string and if it is in the list of defaults specified above, if so it will load that style. If it isn&#8217;t any of those it checks if you have used one of the special options, this loads the style you specified in the plugin settings in the admin control panel for either the top or the bottom. Then it checks if it isn&#8217;t any of those and you have passed it an array instead, if you have it calls another function which builds a custom widget based on what is specified in the array.

The function which builds the custom widget first checks to see if you&#8217;ve specified size and if you&#8217;ve set it to 32 then it add the option for 32px icons. It then loops through the services key which should be another array with a list of services you want to show. Then it checks to see if you&#8217;ve set a number of prefered icons to fill in and adds those to the widget. It checks to see if you&#8217;ve specified for the more button to be shown. Finally it checks to see if you&#8217;ve specified the counter, if you&#8217;ve set counter to true it shows the default or if you&#8217;ve set the counter to a string it uses that string as the counter style.
