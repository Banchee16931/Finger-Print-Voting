<div class="spaced-container">
    <h1>Style Example Page</h1>
    <p>
        This page shows a ton of different styling options available.
    </p>
    <h2>Colours</h2>
    <div class="colour-div">--primary: <div style="background-color: var(--primary);"/></div>
    <div class="colour-div">--secondary: <div style="background-color: var(--secondary);"/></div>
    <div class="colour-div">--background: <div style="background-color: var(--background);"/></div>
    <div class="colour-div">--background-header: <div style="background-color: var(--background-header);"/></div>
    <div class="colour-div">--background-footer: <div style="background-color: var(--background-footer);"/></div>
    <div class="colour-div">--text: <div style="background-color: var(--text);"/></div>
    <div class="colour-div">--text-inverted: <div style="background-color: var(--text-inverted);"/></div>
    
    <h2>Fonts</h2>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 100;">
        Roboto 100
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 100; font-style: italic;">
        Roboto 100 Italic
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 300;">
        Roboto 300
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 300; font-style: italic;">
        Roboto 300 Italic
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 400;">
        Roboto 400
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 400; font-style: italic;">
        Roboto 400 Italic
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 500;">
        Roboto 500
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 500; font-style: italic;">
        Roboto 500 Italic
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 700;">
        Roboto 700
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 700; font-style: italic;">
        Roboto 700 Italic
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 900;">
        Roboto 900
    </p>
    <p style="font-family: 'Roboto', sans-serif; font-weight: 900; font-style: italic;">
        Roboto 900 Italic
    </p>

    <h2>Headings</h2>
    <h1>Heading 1</h1>
    <p>Under Text</p>
    <h2>Heading 2</h2>
    <p>Under Text</p>
    <h3>Heading 3</h3>
    <p>Under Text</p>
    <h4>Heading 4</h4>
    <p>Under Text</p>
    <h5>Heading 5</h5>
    <p>Under Text</p>
    <h6>Heading 6</h6>
    <p>Under Text</p>

    <h2>Small Headings</h2>
    <h1 class="small">Heading 1</h1>
    <p>Under Text</p>
    <h2 class="small">Heading 2</h2>
    <p>Under Text</p>
    <h3 class="small">Heading 3</h3>
    <p>Under Text</p>
    <h4 class="small">Heading 4</h4>
    <p>Under Text</p>
    <h5 class="small">Heading 5</h5>
    <p>Under Text</p>
    <h6 class="small">Heading 6</h6>
    <p>Under Text</p>

    <h2>Paragraphs</h2>
    <h6>Normal</h6>
    <p>I am a normal paragraph <br/> Here me roar!</p>
    <h6>Lead</h6>
    <p class="lead">I am a lead paragraph <br/> Here me roar!</p>
    <h6>Small</h6>
    <p class="small">I am a small paragraph <br/> Here me roar!</p>
    
    <h2>Unordered List</h2>
    <ul>
        <li>I am an element</li>
        <li>I am an element</li>
        <ul>
            <li>I am a nested element</li>
            <li>I am a nested element</li>
        </ul>
        <li>I am an element</li>
    </ul>

    <h2>Ordered List</h2>
    <ol>
        <li>I am an element</li>
        <li>I am an element</li>
        <ol>
            <li>I am a nested element</li>
            <li>I am a nested element</li>
        </ol>
        <li>I am an element</li>
    </ol>

    <h2>Link</h2>
    <a href="/styling">I am a link</a>

    <h2>Text Positioning</h2>
    <p>Normal Text (Left by Default)</p>
    <p class="text-center">Text With "text-center" Class (Center)</p>
    <p class="text-right">Text With "text-right" Class (Right)</p>

    <div style="height: 300px;"/>
</div>

<style lang="scss">
    .colour-div {
        width: 200px;
        display: flex;
        align-items: center;
        div {
            margin-left: auto;
            width: 15px; 
            height: 15px; 
            border: 2px solid black;
        }
    }
</style>