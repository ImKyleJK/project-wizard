# 🧙 project-wizard // templates
Community created templates are what allows project wizard to grow and and give more value to the project.

## Creating a template

1. First find the correct directory to create your template there are 3 options [`Langauge`](/Langauge), [`Frameworks`](/frameworks), [`Stack`](/stack).
2. Now you are in the right place create a `.md` file and name it according to the template.
3. Using the template below create your project template and create a pull request. (Dont forget to add the tag `template`)

<pre><code>name | description | languages | Author
--- | --- | ---
{Template name} | {Template description} | {Languages used in template} | {Your github handle}

# {File name + path}
```
your code goes here ...
```

command name | command | params
--- | --- | ---
{Command name} | {Command} | {Optional params}
</pre></code>

## Example template

<pre><code>name | description | languages
--- | --- | ---
Nodejs | Hello world in nodejs | javascript | NotReeceHarris

# index.js
```
console.log('Hello World')
```

command name | command | params
--- | --- | ---
Setup node package manager | npm init |
</pre></code>
