---
title: How to query a database with AJAX and display as a tooltip
author: Jacob Tomlinson
layout: post
permalink: /2012/09/11/ajax-database-tooltip/
category: Web Development
thumbnail: jquery
tags:
  - AJAX
  - Guide
  - JavaScript
  - jQuery
  - MySQL
  - PHP
---
This post began as an answer on Stack Overflow to a question on &#8216;<a href="http://stackoverflow.com/questions/12291654/netflix-style-tooltip-with-ajax-jquery-php-and-mysql/12366841#12366841" target="_blank">How to query a database with AJAX and display as a tooltip</a>&#8216;. I have put the answer here for future reference.

The original question asked for more information on the steps :

*   User clicks on a link
*   Variable passed to external file using AJAX
*   External file uses variable to query MySQL database
*   Query result appears in tooltip


and so I have turned each step into a heading and added content to each to describe the process.


### User clicks on a link


So the first thing you need is a bit of front end code which tells a link to make an ajax request and not an actual link. I would recommend that you use jQuery to do the hard work for you so you will first need to include the jQuery script in your page if you haven&#8217;t already, to do this add this line to the head of your HTML file.

```html
<script src="http://code.jquery.com/jquery-latest.min.js">
```

Next you will need to give the link a class so for example your link may look like

```html
<a href="ajax.php?var=value" class="ajax_link">Link</a>
```

You will see that the link will be to the ajax page which is going to provide the data, you are passing a variable on to that page called var via GET which we will come to later and you also have the class set to "ajax_link" which we will use to reference the link from jQuery.

Now you will need to write some jQuery to handle the link. To do this you will need to open up a `<script language="javascript">` tag at the end of your body and put some code along the lines of:

```javascript
$(".ajax_link").click(function(e) {

    e.preventDefault(); //Stops link from changing the page

    var link = $(this).attr('href'); //Gets link url

    $.ajax({ //Make the ajax request
      url: link,
      cache: false
    }).done(function( html ) { //On complete run tooltip code

        //Display tooltip code goes here, returned text is variable html

    });
});
```

I will come back to the tooltip code later on.

### Variable passed to external file using AJAX
Now we are going to need to create the PHP file ajax.php, this file will need to connect to the database and make any queries necessary. First I will start the file and get the variable from the GET. See the code below.

```php
$var = mysql_real_escape_string($_GET['var']);</pre>
```

This will create the variable $var and it will contain the value &#8216;value&#8217; as was set in the above HTML link. You will probably want to escape the value, do some validation to protect yourself from injection but I will not cover that here.  

**UPDATE:** Added most basic string escape function to stop people from copying and pasting and opening themselves up to injection but I would recommend seeing <a title="Stack Overflow MySQL Injection" href="http://stackoverflow.com/a/7528395/1003288" target="_blank">this</a> stack overflow answer for more information.

### External file uses variable to query MySQL database  
Next you will want to connect to and query the database and then echo what ever it is that you have returned. Here is some code for returning the value of column &#8216;information&#8217; from a table called &#8216;data_table&#8217; where some value equals your $var.

```php
//connection to the database
$dbhandle = mysql_connect($hostname, $username, $password)
  or die("Unable to connect to MySQL");

//select a database to work with
$selected = mysql_select_db("examples",$dbhandle)
  or die("Could not select examples");

//execute the SQL query and return records
$result = mysql_query("SELECT information FROM data_table WHERE value='$var'");
//fetch tha data from the database
while ($row = mysql_fetch_array($result)) {
   echo $row{'information'};
}

//close the connection
mysql_close($dbhandle);</pre>
```

For more information on querying a database please see <a href="http://webcheatsheet.com/php/connect_mysql_database.php" target="_blank">this tutorial</a>.

### Query result appears in tooltip
Once you have both of these pages set up the first page will send some value to the ajax page, the ajax page will echo the corresponding information from the database which will be received by the first page in the JavaScript variable **html**.

The final stage of this is now to display the returned text as a tooltip. First it may be a good idea to test that your pages are working and one easy way to do this is to display an alert with the returned information in (you could also log to the console). You would do this by adding some code to the ajax .done function like so:

```javascript
.done(function( html ) {
    alert("text: " + html);
});
```

Once you have done this you should be able to open up the page, click the link and you should see an alert box pop up showing the required text. If there is nothing after the &#8216;text:&#8217; in the alert then something has gone wrong and you may want to check the code is correct. If it has succeeded then continue on to displaying the tooltip.

There are many available jQuery tooltip plugins one of which is called <a href="http://craigsworks.com/projects/qtip2/" target="_blank">qTip</a>. This will allow you to create the tooltips and display them when clicking on them. You will need to download the JavaScript file and related CSS for the tooltips and then add some code into the .done function to create and display the tooltip.

To add qTip to your page place visit the qTip site and download the latest version of the js and the css. Then you must reference the files in the head like this:

```html
<link type="text/css" rel="stylesheet" href="jquery.qtip.min.css" />
```

and the js like this:

```html
<script type="text/javascript" src="jquery.qtip.min.js"></script>
```

This assumes that the css and js files are in the same folder as your index page, if they aren&#8217;t please amend these links to point to where you have put those files.

Now all you need to add to the .done function is the following:

```javascript
$(this).qtip({
    content: {
        text: html
    }
});
$(this).qtip('toggle', true);
```

This code will create a qTip containing whatever was returned from your AJAX request. Then once the tooltip has been created it will toggle it to show the tooltip initially once clicked. Then the tooltip will show/hide when you hover on/off.

For more information on styling a qTip see their website.
