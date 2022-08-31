---
title: "A guide on HTML reprs and ipywidgets for non-web developers"
date: 2021-06-07T00:00:00+00:00
draft: true
author: "Jacob Tomlinson"
categories:
  - blog
tags:
  - Python
  - HTML
  - Jupyter
  - ipywidgets
---

Jupyter allows you to build rich and interactive HTML representations for your objects when folks inspect them within their notebook. However building these can be daunting for folks that do not consider themselves web developers. In this post I intend to give a detailed overview on what is possible and how you can add awesome reprs and widgets to your objects.

## Python reprs

Let's start at the beginning with Python reprs.

A "repr" is a string representation of a Python object designed to help folks explore and debug their Python code. All objects in Python come with these by default, you've probably seen them, they are little angle bracket wrapped descriptions of the object itself.

If you create your own custom object and have a look at the repr in ipython you will see the package path of the object and a pointer to that instance.

```python
In [1]: class MyAwesomeObject(object):
   ...:     """My awesome object."""
   ...:

In [2]: mao = MyAwesomeObject()

In [3]: mao
Out[3]: <__main__.MyAwesomeObject at 0x7f87d8b161f0>
```

Python kindly generated this for us, but we could choose to specify our own `__repr__` method on our object so that we can generate our own.

```python
In [1]: class MyAwesomeObject(object):
   ...:     """My awesome object."""
   ...:
   ...:     def __repr__(self):
   ...:         return f"<{self.__class__.__name__} How awesome is this?>"
   ...:

In [2]: mao = MyAwesomeObject()

In [3]: mao
Out[3]: <MyAwesomeObject How awesome is this?>
```

This can be extremely useful if our object has some attributes that we want to convey.

```python
In [1]: class MyAwesomeObject(object):
   ...:     """My awesome object."""
   ...:
   ...:     def __init__(self, foo, bar, baz):
   ...:         self.foo = foo
   ...:         self.bar = bar
   ...:         self.baz = baz
   ...:
   ...:     def __repr__(self):
   ...:         return f"<{self.__class__.__name__} foo='{self.foo}' bar='{self.bar}' baz='{self.baz}'>"
   ...:

In [2]: mao = MyAwesomeObject("abc", "def", "ghi")

In [3]: mao
Out[3]: <MyAwesomeObject foo='abc' bar='def' baz='ghi'>
```

These reprs are going to be primarily used when debugging so giving the user some information about the object is going to help them way more than telling them the memory address of the object.

## Jupyter HTML reprs

If we jump over to Jupyter we can see that the same code also shows our repr inline in our notebook.

![Python repr displayed in Jupyter Lab](https://i.imgur.com/V94H5mN.png)

But Jupyter is a web application. The plain text repr we are seeing on the page has been rendered as HTML in order for us to view it.

If we inspect the element on the page using our browser's built in debugging tools we can see that it has been wrapped in a `<pre>` element which ensures that whitespace including new lines is preserved and that has been wrapped in a`<div>` element with some classes attached.

![Inspecting a plain text repr in the browser](https://i.imgur.com/Ur1NO3t.png)

Jupyter is awesome because it allows us to also define a `_repr_html_` method which is used to generate that HTML. This allows us to do something richer, more bespoke and more stylized for Jupyter users.

```python
class MyAwesomeObject(object):
    """My awesome object."""

    def __init__(self, foo, bar, baz):
        self.foo = foo
        self.bar = bar
        self.baz = baz

    def __repr__(self):
        return f"<{self.__class__.__name__} foo='{self.foo}' bar='{self.bar}' baz='{self.baz}'>"

    def _repr_html_(self):
        return f"""
        <h3>{self.__class__.__name__}</h3>
        <ul>
          <li>foo='{self.foo}'</li>
          <li>bar='{self.bar}'</li>
          <li>baz='{self.baz}'</li>
        </ul>
        """
```

![Our first HTML repr](https://i.imgur.com/646GZPE.png)

Now instead of printing our regular repr to the screen we can see our rich text HTML repr which puts the class name in a `<h3>` heading element and shows the attributes as a bullet point list.

We can also inspect this in our browser tools and see that there is no `<pre>` block, instead it has been replaced with the contents of our HTML method.

![Inspecting the HTML repr](https://i.imgur.com/URZvwsM.png)

Awesome! Now let's make some nice HTML reprs!

## How can we best leverage HTML?

I'm going to take a guess that the average reader of this post is aware of HTML and that it is used in building web pages, but does not consider themselves a web developer. Web pages are typically built with HTML which represents the contents of the page, CSS which styles that content and JavaScript which makes that content interactive and dynamic.

We are going to ignore JavaScript for now. We will come onto interactivity later with `ipywidgets`. So let's talk about how we can use HTML and CSS to make useful representations of our objects.

Typically when you create a website with HTML and CSS you put the HTML in one file and the CSS in another and then link the two together via a piece of metadata in the HTML. However for our HTML reprs we are just providing Jupyter with a snippet of code, not a whole page, so all of our styling needs to happen inline.

Also much of what we are going to build here will be done within multi-line Python strings. So we really need to keep the code as lean and readable as possible otherwise things will get messy. To do this we need to really lean on traditional HTML.

### Styling with CSS

We should still use CSS to style our content. For simple HTML reprs we may even want to include our CSS inline, but be aware that this will get messy fast. If we only want to add a couple of rules to a few elements then that's probably ok, but when you start adding more and more it'll begin to look like this.

```html
<h3 style="color: red; font-size: 23px; font-weight: 500; margin-top: 2px; margin-bottom: 5px;">{self.__class__.__name__}</h3>
```

To style our content we are going to place a block of CSS above our HTML, all major browsers support blocks of CSS throughout the body like this but it feels slightly odd for folks who are familiar with web design.

```python
    def _repr_html_(self):
        style = """
        <style type="text/css">
        h3 { color: red; }
        </style>
        """

        body = f"""
        <h3 class="mao-title">{self.__class__.__name__}</h3>
        <ul>
          <li>foo='{self.foo}'</li>
          <li>bar='{self.bar}'</li>
          <li>baz='{self.baz}'</li>
        </ul>
        """

        return style + body
```

For example we can change the text colour of the `<h3>` element and it will display like this.

![HTML repr with a red title](https://i.imgur.com/mD7MSPO.png)

It is important to be aware that this CSS applies to the whole page, so we just made all `<h3>` objects red. Therefore it is much better to use class selectors and class names that are unique to your repr.

Classes in CSS are a common way to style a similar set of objects. These are specified in your HTML by the `class="foo"` attribute and then matched in your CSS with a `.foo { ... }` selector.

```python
    def _repr_html_(self):
        style = """
        <style type="text/css">
        .mao-title { color: red; }
        </style>
        """

        body = f"""
        <h3 class="mao-title">{self.__class__.__name__}</h3>
        <ul>
          <li>foo='{self.foo}'</li>
          <li>bar='{self.bar}'</li>
          <li>baz='{self.baz}'</li>
        </ul>
        """

        return style + body
```

Putting the two sections into separate variables like this kinda starts to feel like more traditional web dev, CSS in one place, HTML in another. We also made use of a f-string to format our HTML string and if we included our CSS in there we would need to escape our curly braces, so this avoids that problem.

### HTML elements

Let's talk about some good HTML elements to use in your repr. Personally I think an HTML repr should be pretty compact and information dense so let's look at some useful elements for achieving that.

#### Tables

A great way to show a lot of information is with a table, and Jupyter also styles these really nicely for us out of the box.

Remember those classes that Jupyter put on the surrounding `<div>` earlier? Those are some nice base styles which give all our elements a consistent look. There are a few things you might want to tweak though, for example tables are not full width by default and cells are right aligned, so let's tweak that in our CSS.

```python
    def _repr_html_(self):
        style = """
        <style type="text/css">
          .mao-table {
            width: 100%;
          }
          .mao-table td {
            text-align: left !important;
          }
        </style>
        """

        body = f"""
        <h3 class="mao-title">{self.__class__.__name__}</h3>
        <table class="mao-table">
        <tr>
          <td>foo='{self.foo}'</td>
          <td>bar='{self.bar}'</td>
        </tr>
          <td>baz='{self.baz}'</td>
          <td></td>
        </tr>
        </ul>
        """

        return style + body
```

![HTML repr with a table](https://i.imgur.com/zefI0P2.png)

Here we have switched our list to a table using `<table>`, `<tr>` (table row) and `<td>` (table cell) elements. We've also used an `!important` flag in our CSS to ensure that option is selected. Typically `!important` flags are bad practice in CSS and you should adjust your style hierarchy instead. However in this case we have no control over the style hierarchy, we are limited to the CSS that we have created.

#### Text styling

Next we might want to think more about styling our text. We can of course add CSS to change things like size, colour, etc, but it may keep things leaner to sparingly use HTML styling elements.

```python
    def _repr_html_(self):
        style = """
        <style type="text/css">
          .mao-table {
            width: 100%;
          }
          .mao-table td {
            text-align: left !important;
          }
        </style>
        """

        body = f"""
        <h3 class="mao-title">{self.__class__.__name__}</h3>
        <table class="mao-table">
        <tr>
          <td><b>foo:</b> {self.foo}</td>
          <td><b>bar:</b> {self.bar}</td>
        </tr>
          <td><b>baz:</b> {self.baz}</td>
          <td></td>
        </tr>
        </ul>
        """

        return style + body
```

![Html repr with table and bold text](https://i.imgur.com/sp3UbEO.png)

Here we've wrapped the attribute names in `<b>` elements to bold the text.

There are a bunch of different elements which we could use to modify our text.

- `<b>` Makes text bold
- `<i>` Makes text italic
- `<u>` Makes text underlined
- `<s>` Strikesthrough text
- `<sup>` Makes text supertext
- `<sub>` Makes text subtext
- `<strong>` Makes text bold and adds semantic information that this text is important
- `<em>` Makes text italic and adds semantic information that this text is emphasised

For example here is how these different elements look in Jupyter.

![Preview of different text styles](https://i.imgur.com/CWVQan4.png)

HTML is primarily about providing structure and semantics but historically has elements for styling too. It is more common these days to use CSS for styling and only use HTML for semantics. For our use case though my preference is to keep our repr code as lean as possible and `<b>` takes up much less room than a CSS rule to increase the font weight.

#### Layout

The next thing to think about is layout. This covers order of elements on the page and their size, padding and margins.

If we look at our `<h3>` element for example we can see that there are it quite a large margin at the top which pushes the text down.

![h3 element with margin](https://i.imgur.com/2BXEoB2.png)

This would make sense if it were a true level three title element in a piece of text, but in this context we are using that element for size and weight and we do not want that extra padding. Let's change that margin to something smaller.

```python
    def _repr_html_(self):
        style = """
        <style type="text/css">
          .mao-title {
            margin-top: 0.25em !important;
          }
          .mao-table {
            width: 100%;
          }
          .mao-table td {
            text-align: left !important;
          }
        </style>
        """

        body = f"""
        <h3 class="mao-title">{self.__class__.__name__}</h3>
        <table class="mao-table">
        <tr>
          <td><b>foo:</b> {self.foo}</td>
          <td><b>bar:</b> {self.bar}</td>
        </tr>
          <td><b>baz:</b> {self.baz}</td>
          <td></td>
        </tr>
        </ul>
        """

        return style + body
```

![HTMl repr with smaller top margin](https://i.imgur.com/QmT58g1.png)

Now the padding is much smaller and things align better.

You may have noticed we used a unit called `em`. CSS has a bunch of units you can use including `px` for pixels, `%` for percentage of the parent element, `em` where `1em` is equal to the font size of the current element and `rem` where `1rem` is equal to the default font size on the page.

Using `em` and `rem` is nice because if a user increases their font size in their browser or zooms in the padding and margins will also increase.

#### Layout elements

HTML also has a few layout elements which can be useful.

The horizontal rule element `<hr>` creates a horizontal line across the page. This can be useful for breaking up sections.

```html
<p>This text is the main part of the information</p>
<hr />
<sup>This is a footnote.</sup>
```

![Example using hr element](https://i.imgur.com/EfCMqFT.png)

Web browsers also try to collapse whitespace so even if your HTML includes line breaks they will likely not be honored.

```html
The quick brown fox
jumps over the lazy dog.
```

![Sentence rendered on one line](https://i.imgur.com/TSzPvuF.png)

Despite there being a break in the text the browser renders this sentence on one line. There are a few solutions to this.

We can add a line break `<br />` element.

```html
The quick brown fox<br />
jumps over the lazy dog.
```

![Sentence rendered with a line break](https://i.imgur.com/tXWD6Cq.png)

We could also wrap our text in a `<pre>` element which tells the browser to preserve all whitespace. However it is worth noting that this element is also commonly used for quotes and to wrap code blocks and the default Jupyter styling adds a margin to the left hand side.

```html
<pre>
The quick brown fox
jumps over the lazy dog.
</pre>
```

![Sentence rendered with a line break and left margin](https://i.imgur.com/8yY3szi.png)

Lastly we could use `<p>` elements. Semantically a `<p>` section represents a paragraph, so it shouldn't really be used for breaking up sentences but instead ensuring there is correct spacing between paragraphs. Commonly this adds the line break and also some margins.

```html
<p>
The quick brown fox jumps over the lazy dog.
</p>
<p>
A mad boxer shot a quick, gloved jab to the jaw of his dizzy opponent.
</p>
```

![Two sentences broken up into paragraphs](https://i.imgur.com/cI7FUNG.png)

### What about images?

Images are pretty common things to include on web pages, what about those?

My opinion here is that our HTML repr should be completely self contained. It should not rely on any external assets. So while it might be tempting to upload images or icons to a CDN like Imgur and use regular `<img src="" />` elements I would advise against it.

So what does that leave us with?

#### Base64 encoded PNGs

If the image you want to use is very small you could create a base64 encoded version and include it inline.

Base64 encoding an image means that the image will be converted to a string of text which can be passed entirely in HTML. For a quick example I've gone to Google Slides and drawn a little logo for our object.

![Logo design in google Slides](https://i.imgur.com/dvaaDmh.png)

Nothing fancy here, just some coloured text in a fun font. Then I can take a screenshot of this and run it through a [base64 encoder](https://www.gieson.com/Library/projects/utilities/base64-image/). This gives me some HTML that I could use in my HTML repr.

```html
<img alt="my image" width=100 height=84 src="data:image/png;base64,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" />
```

This element is pretty big, but at the same time it's not insanely huge either. We could include this directly in our repr string.

```python
class MyAwesomeObject(object):
    """My awesome object."""

    def __init__(self, foo, bar, baz):
        self.foo = foo
        self.bar = bar
        self.baz = baz

    def __repr__(self):
        return f"<{self.__class__.__name__} foo='{self.foo}' bar='{self.bar}' baz='{self.baz}'>"

    def _repr_html_(self):
        style = """
        <style type="text/css">
          .mao-title {
            margin-top: 0.25em !important;
            margin-left: 3em !important;
          }
          .mao-logo {
            position: absolute;
            height: 3em !important;
          }
          .mao-table {
            width: 100%;
          }
          .mao-table td {
            text-align: left !important;
          }
        </style>
        """

        body = f"""
        <img alt="my image" class="mao-logo" src="data:image/png;base64,..." />
        <h3 class="mao-title">{self.__class__.__name__}</h3>
        <table class="mao-table">
        <tr>
          <td><b>foo:</b> {self.foo}</td>
          <td><b>bar:</b> {self.bar}</td>
        </tr>
          <td><b>baz:</b> {self.baz}</td>
          <td></td>
        </tr>
        </ul>
        """

        return style + body
```

![HTML repr with base64 png logo](https://i.imgur.com/n4kGUjA.png)

Despite being easy to work with there are a number of downsides to using PNGs like this.

- You can't edit them easily.
- You can't rescale them and retain clarity.
- They are big and messy.
- They take up a large amount of space in your code.

#### Inline SVG images

My preference for images in HTML reprs is to use SVG images inline in your code.

SVGs are a vector graphic standard which complements the HTML standard. This means instead of a base64 encoded raster image you have a programmatic description of how to draw the image.

For drawing SVGs I like to use Figma. It's a free design tool which is pretty beginner friendly. Here I've quickly created the same piece of text that I did in Google Slides with the same font and colours.

![Drawign in Figma](https://i.imgur.com/YISWCTX.png)

But in Figma I can right click on the text and copy it as an SVG.

![Copy as SVG](https://i.imgur.com/LOfm6C0.png)

```svg
<svg width="68" height="52" viewBox="0 0 68 52" fill="none" xmlns="http://www.w3.org/2000/svg">
<path d="M19.9613 51.472C19.4066 51.472 19.0226 51.3227 18.8093 51.024C18.6386 50.768 18.5533 50.32 18.5533 49.68V43.856C18.5533 40.9547 18.4679 37.4347 18.2973 33.296C18.2546 32.528 18.2333 31.4187 18.2333 29.968C18.1053 26.9387 18.0413 24.5493 18.0413 22.8L18.1053 19.216C18.1479 18.1493 18.1693 16.9547 18.1693 15.632C18.1693 14.48 18.1479 13.456 18.1053 12.56C18.0626 11.92 18.0199 11.3867 17.9773 10.96C17.9773 10.5333 17.9559 10.192 17.9133 9.936L17.7853 10.384C17.3159 13.4987 16.6119 17.5947 15.6733 22.672C14.7346 27.7493 13.9666 31.824 13.3693 34.896C12.7719 37.6267 12.3879 39.568 12.2173 40.72C12.1319 41.2747 11.9399 41.7227 11.6413 42.064C11.3853 42.4053 11.0866 42.576 10.7453 42.576C10.4039 42.576 10.1053 42.4053 9.84925 42.064C9.59325 41.7227 9.44392 41.232 9.40125 40.592C9.10258 38.544 8.01458 32.8267 6.13725 23.44C4.25992 14.0533 3.21458 9.18933 3.00125 8.848C2.87325 12.688 2.80925 16.9973 2.80925 21.776C2.80925 25.5307 2.83058 28.9227 2.87325 31.952C2.91592 34.9813 2.93725 38.3733 2.93725 42.128C2.93725 42.768 3.00125 43.7707 3.12925 45.136L3.25725 48.208C3.25725 49.0187 3.10792 49.5733 2.80925 49.872C1.57192 49.872 0.95325 48.7627 0.95325 46.544V35.856L1.01725 29.392C1.05992 27.472 1.08125 25.296 1.08125 22.864C1.08125 20.048 1.10258 17.5307 1.14525 15.312L1.27325 7.696V2.064C1.27325 1.85067 1.35858 1.65867 1.52925 1.488C1.74258 1.31733 1.95592 1.232 2.16925 1.232C2.51058 1.232 2.74525 1.40267 2.87325 1.744C3.25725 3.024 4.21725 6.8 5.75325 13.072C7.28925 19.344 8.44125 24.1867 9.20925 27.6C9.59325 29.6053 10.2119 33.1467 11.0653 38.224L11.5133 35.728C11.6413 34.96 11.8333 33.8933 12.0893 32.528C12.6013 29.2 12.9213 27.3227 13.0493 26.896L15.9933 11.216L16.5693 7.888L17.2733 4.304L17.5933 2.384C17.6359 1.95733 17.7639 1.63733 17.9773 1.424C18.1906 1.168 18.4253 1.04 18.6813 1.04C18.8946 1.04 19.0866 1.14666 19.2573 1.36C19.4279 1.57333 19.5133 1.872 19.5133 2.256V8.592C19.5133 11.792 19.7053 16.6347 20.0893 23.12C20.3879 29.0933 20.5373 33.8933 20.5373 37.52C20.5373 38.672 20.6013 40.4427 20.7293 42.832C20.8146 45.904 20.8573 48.0373 20.8573 49.232C20.8573 50.1707 20.7933 50.7893 20.6653 51.088C20.4946 51.344 20.2599 51.472 19.9613 51.472Z" fill="#CC4125"/>
<path d="M42.4258 50C42.0844 50 41.7858 49.8933 41.5298 49.68C41.3164 49.4667 41.2098 49.2107 41.2098 48.912V40.528C41.2098 40.016 40.1218 39.5467 37.9458 39.12C35.7698 38.6933 33.9778 38.48 32.5698 38.48C32.4844 38.48 32.3351 38.8 32.1218 39.44C31.9511 40.08 31.8231 40.5493 31.7378 40.848C31.1831 43.0667 30.6284 45.4347 30.0738 47.952L29.7538 49.36C29.6684 49.7013 29.4551 49.872 29.1138 49.872C28.7724 49.872 28.4524 49.7653 28.1538 49.552C27.8978 49.3387 27.7698 49.0827 27.7698 48.784L27.8338 48.656V48.592L28.5378 45.392C29.0498 43.3867 29.4978 41.424 29.8818 39.504C30.3084 37.328 30.6284 35.92 30.8418 35.28C31.1831 33.872 31.8658 30.352 32.8898 24.72L33.2738 22.608C34.3831 16.0373 35.1511 11.8987 35.5778 10.192C35.7911 9.168 36.1751 7.65333 36.7298 5.648C36.9431 4.79466 37.2418 3.94133 37.6258 3.088C38.0098 2.36266 38.3298 2 38.5858 2C39.2684 2 39.7591 2.21333 40.0578 2.64C40.2711 3.024 40.5911 5.136 41.0178 8.976C41.4871 12.816 41.8924 16.6987 42.2338 20.624L43.0658 31.248C43.0658 33.0827 43.1938 35.856 43.4498 39.568C43.7058 43.3227 43.8338 46.1173 43.8338 47.952V48.208C43.8338 48.6773 43.7698 49.104 43.6418 49.488C43.4711 49.8293 43.0658 50 42.4258 50ZM39.8658 37.648C40.7191 37.648 41.1458 37.4347 41.1458 37.008C41.1458 34.3627 40.9964 30.7573 40.6978 26.192C40.3991 21.584 40.0791 17.5307 39.7378 14.032C39.6524 12.5387 39.4818 10.896 39.2258 9.104L38.9698 6.608C38.9698 6.65066 38.9271 6.416 38.8418 5.904L38.7138 5.136C38.3724 6.07467 38.0311 7.376 37.6898 9.04C37.3911 10.704 37.0711 12.6667 36.7298 14.928C35.7058 21.2853 34.8311 26.064 34.1058 29.264L33.5298 31.44C32.8898 33.5733 32.5698 35.1947 32.5698 36.304V36.688C34.4471 37.2853 36.6018 37.584 39.0338 37.584H39.4178C39.5458 37.6267 39.6951 37.648 39.8658 37.648Z" fill="#93C47D"/>
<path d="M60.9563 51.344C55.0256 51.344 52.0603 41.0827 52.0603 20.56V18C52.0603 14.416 52.2736 11.3013 52.7003 8.656C52.9563 6.992 53.2976 5.62666 53.7243 4.56C54.1936 3.49333 54.8123 2.64 55.5803 2C56.3056 1.27466 57.2443 0.911999 58.3963 0.911999C59.4203 0.911999 60.3589 1.12533 61.2123 1.552C62.0656 1.936 62.7696 2.448 63.3243 3.088C63.7936 3.68533 64.2203 4.47466 64.6042 5.456C64.9883 6.43733 65.2869 7.376 65.5003 8.272L66.0123 11.28C67.0363 18.576 67.5483 25.04 67.5483 30.672V30.864C67.5483 32.9973 67.5056 34.96 67.4203 36.752C67.3349 38.8853 67.1003 41.1253 66.7163 43.472C66.5456 44.624 66.2896 45.6267 65.9483 46.48C65.6496 47.3333 65.2443 48.1867 64.7323 49.04C63.8363 50.576 62.5776 51.344 60.9563 51.344ZM61.0843 49.552C61.7669 49.552 62.3856 49.168 62.9403 48.4C63.4949 47.5893 63.9216 46.48 64.2203 45.072C64.9029 42.0853 65.3296 39.1627 65.5003 36.304C65.5856 34 65.6283 31.76 65.6283 29.584V29.328L65.5643 22.928C65.4789 20.624 65.2229 17.6373 64.7963 13.968C64.4549 11.3653 63.7509 8.69867 62.6843 5.968C62.2576 4.90133 61.6816 4.06933 60.9563 3.472C60.2736 2.87466 59.5269 2.576 58.7163 2.576C57.9909 2.576 57.3509 2.91733 56.7963 3.6C56.2416 4.28266 55.7936 5.24266 55.4523 6.48C54.6843 9.33867 54.2576 11.984 54.1723 14.416C54.0443 16.5067 53.9803 18.6613 53.9803 20.88V24.016C53.9803 26.8747 54.2576 30.5013 54.8123 34.896C55.1536 37.6693 55.5803 40.08 56.0923 42.128C56.6043 44.176 57.2869 45.904 58.1403 47.312C58.9936 48.8053 59.9749 49.552 61.0843 49.552Z" fill="#6D9EEB"/>
</svg>
```

This SVG is pretty big. I specifically chose it because the text has complex lines which curve in certain ways which makes the SVG larger. But if you read through this SVG it should make more sense than the base64 encoded image.

We can see that it is wrapped in a top level `<svg>` element with some size and meta attributes. Then within the SVG there are three `<path>` elements each with a `d` (data) value and a `fill` colour. This data describes the path that makes up the shape, each number being a coordinate to move the virtual pen to. We probably don't want to edit that by hand, but we could definitely change the colour of each letter if we wanted to.

SVG really shines when you want draw simple shapes though. Let's head back to Figma and draw a blue circle with a black outline.

![Figma drawing of a blue circle with a black outline](https://i.imgur.com/JXWghed.png)

If we copy the SVG for this element we see it is much simpler.

```svg
<svg width="115" height="115" viewBox="0 0 115 115" fill="none" xmlns="http://www.w3.org/2000/svg">
<circle cx="57.5" cy="57.5" r="52.5" fill="#6D9EEB" stroke="black" stroke-width="10"/>
</svg>
```

Our SVG contains one `<circle>` element with `cx` and `cy` coordinates describing the origin of the circle and an `r` element describing the radius. It a `fill` attribute to describe the colour and `stroke` and `stroke-width` attributes to describe that outline. This is much easier for humans to understand and if we wanted to we could tweak this code to change things without going back into Figma.

Using simple SVG shapes like this is really powerful for adding interesting and unique looking visuals in a small amount of code.

#### Drawing with divs

The last image type I want to touch on is styling HTML elements to make shapes. You can do this with any element but it is common to use the `<div>` division element.

We can recreate our SVG circle using CSS.

```python
    def _repr_html_(self):
        style = """
        <style type="text/css">
          .mao-title {
            margin-top: 0.25em !important;
            margin-left: 2em !important;
          }
          .mao-logo {
            position: absolute;
            margin-top: 0.2em;
            width: 1.5em;
            height: 1.5em;
            background-color: #6D9EEB;
            border-radius: 2em;
            border: 3px solid black;
          }
          .mao-table {
            width: 100%;
          }
          .mao-table td {
            text-align: left !important;
          }
        </style>
        """

        body = f"""
        <div class="mao-logo"></div>
        <h3 class="mao-title">{self.__class__.__name__}</h3>
        <table class="mao-table">
        <tr>
          <td><b>foo:</b> {self.foo}</td>
          <td><b>bar:</b> {self.bar}</td>
        </tr>
          <td><b>baz:</b> {self.baz}</td>
          <td></td>
        </tr>
        </ul>
        """

        return style + body
```

![HTML repr with styled div icon](https://i.imgur.com/krpJCPd.png)