---
title: Quick Tip - em vs rem
author: Jacob Tomlinson
layout: post
category: Quick Tip
thumbnail: code
tags:
- quick tips
- em
- rem
excerpt: 'Do you know the difference between em and rem?'
---

`em` and `rem` are used in CSS to set a size value relative to a `font-size`. This is useful in many situations such as increasing the font size relatively across your whole website by changing one value or adding padding which is larger or smaller depending on the font size.

### em

`em` means relative to the font size of the parent DOM element. Therefore if you nest divs with a `font-size` of `0.75em` the font will get increasingly smaller inside each nested div.

```
<style>
html {
  font-size: 1em;
}

div {
  font-size: 0.75em;
}
</style>

<div>
  Small
  <div>
    Smaller
    <div>
      Smallest
    </div>
  </div>
</div>
```

### rem

`rem` means relative to the `font-size` of the root `html` element. Therefore nesting divs with a `font-size` of `0.75rem` will keep the same font size despite the nesting level.

```
<style>
html {
  font-size: 1em;
}

div {
  font-size: 0.75rem;
}
</style>

<div>
  Small
  <div>
    Not Smaller
    <div>
      Still the same size
    </div>
  </div>
</div>
```
