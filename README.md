# tmplr

tmplr is a cli tool to quickly create files from templates.

## Usage

Assuming we are in an empty directory:

```bash
$ tmplr src/index.php
```

This will create the directory "src" and an empty index.php file because we have not created a template just yet.

Lets create one!

First lets go into our template directory:

```bash
$ cd $(tmplr --template-dir)
```

next lets create a template file for PHP classes and PHP enums

```bash
$ hx [class].php
```

```php
---
name: PHP Class Template
vars:
  - name: namespace
    prompt: Namespace for this class
---
<?php

namespace {{namespace}};

class {{class}} {
    public function __constructor() {
    }
}
```

and...

```bash
$ hx [enum].php
```

```php
---
name: PHP Enum Template
vars:
  - name: namespace
    prompt: Namespace for this class
---
<?php

namespace {{namespace}};

enum {{enum}} {
}
```

Short explanation of the template syntax:

There is a frontmatter where you can define a name and variables the user has to enter when they are creating
the template. And after that there is the actual template that will be later parsed by Gos template/text
templating engine.

lets get back to the project dir and create a new file:

```bash
$ cd ~/to/my/project-dir
$ tmplr src/controller/IndexController.php

This filename matches multiple templates, please select one:

    - [x] PHP Class Template ([class].php)
    - [ ] PHP Enum Template  ( [enum].php)

Namespace for this class:
> App\Controller

Template "src/controller/IndexController.php" created.

$ cat src/controller/IndexController.php
<?php

namespace App\Controller;

class IndexController {
    public function __constructor() {
    }
}
```

tmplr matches the filename against all files in its template directory in this case "IndexController.php" matches "\[class\].php" and "\[enum\].php"

The \[...\] part is variable and will as you can see above also be exported to the template.

You could also create a template like this "\[controllerName\]Controller.php" which would also match for "IndexController" to make even more precise templates

## Name

tmplr comes from a mixture of template and Templar and is also supposed to be pronounced "Templar".

## License

GNU General Public License v3

![](https://www.gnu.org/graphics/gplv3-127x51.png)
