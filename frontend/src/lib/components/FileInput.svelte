<script lang="ts">
    import Symbol from "./Symbol.svelte";

    // Defines how it would be submitted to a <form> element. This will act as this inputs id
    export let name: string;

    // Text above the file input to describe what it is to the user
    export let label: string | null = null

    // Bind to this to get the files that have been selected
    export let files: FileList | null = null

    // Set this to allow the user to upload more than one file
    export let multiple: boolean = false;

    // Set what files are allowed in the form of a comma seperated list, by either extension or MIME type (e.g. image/png,.jpeg,.jpg)
    export let accept: string = ""

    // Communicate to the user that this input must be filled for the form to be valid
    export let required: boolean = false;

    // Set this to display a spinner and stop user input
    export let working: boolean = false;

    // Message displayed at bottom of input to show when something has gone wrong
    let failureMessage: string | null = null

    // A list of all the acceptance criteria for the input
    $: acceptanceCriteria = accept.split(",")

    // Stops the file from just opening in the browser rather than being given the drop event
    function dragOver (e: DragEvent) {
        e.preventDefault()
    }

    // Checks and the uploaded any files that are dropped onto the input
    function drop (e: DragEvent) {
        failureMessage = null // resets failure message

        e.preventDefault()
        if (e.dataTransfer?.files) { // checks that the data transfer to the input has files
            let newFiles = e.dataTransfer?.files

            if (newFiles.length <= 0) { // stops an empty file list being uploaded
                failureMessage = "no files in upload"
                files = null
                return
            } else if (newFiles.length > 1 && !multiple) { // stops multiple files being uploaded if the multiple option isn't selected
                failureMessage = "too many files"
                files = null
                return
            }

             
            acceptanceCriteria = acceptanceCriteria.filter((item) => {
                if (item.trim() === "") {
                    return false
                }

                return true
            })

            var re = /(?:\.([^.]+))?$/; // checks if a criteria is an extension

            let resetFiles = false
            if (acceptanceCriteria.length > 0) {
                (Array.from(newFiles)).forEach(file => {
                    let valid = false

                    let filesExt = re.exec(file.name)
                    acceptanceCriteria.forEach(criteria => {
                        if (criteria.startsWith(".")
                            && filesExt != null
                            && filesExt.length > 0 
                            && filesExt?.at(0) === criteria) { // validate as an extension
                                valid = true
                        } else if (file?.type === criteria) { // vaidate via file type
                            valid = true
                        }
                    });

                    if (!valid) {
                        failureMessage = `uploaded file (${file.name}) is of invalid type, accepted inputs: ` + accept
                        resetFiles = true
                        return
                    }
                })
            }

            if (resetFiles) {
                files = null
                return
            }
                
            files = newFiles
        }
    }

    // This will add spaces inbetween the values in the accept criteria
    function formatAccept(text: string): string {
        return text.replaceAll(/,\s*/g, ", ")
    }

    // Resets the failureMessage when the data changes
    function change() {
        failureMessage = null
    }
</script>

<!--
    @component
    # FileInput
    FileInput allows someone to upload a file.

    - (required) Use `name` to define how it would be submitted to a <form> element. This will act as this inputs id
    - Use `files` to get the files from the input
    - Use `label` to write some text above the input to describe what it is to the user
    - Use `required` to communicate to the user that this input must be filled for the form to be valid
    - Use `working` to show a spining icon so the user knows that you are processing the input
    - Use `multiple` to allow for multiple files to be uploaded
    - Use `accept` to set what files are allowed in the form of a comma seperated list, by either extension or MIME type (e.g. image/png,.jpeg,.jpg)

    ## Usage
    ```tsx
    <FileInput name="example"/>
    ```

    ### Binding the files
    ```tsx
    <script lang="ts">
        let myFiles: FileList | null;
    </script>
    <FileInput name="example" bind:files={myFiles}/>
    ```
-->
<div class="file-input">
    <small class="label">{label}</small>
    <!-- create a dotted outline arroudn the input to show where files can be dropped -->
    <div class="container" 
    class:invalid={required && (files?.length === 0 || files == undefined)}
    on:drop={e => drop(e)} on:dragover={dragOver} role="form"  {...$$restProps}>
        <!-- innard is used to center the content -->
        <div class="innard">
            <!-- displays upload symbol normally and a spinner if working -->
            {#if !working}
                <Symbol name="upload" color="primary" variant="outlined" style="font-size: 4rem; grid-area: symbol; padding-right: 10px;"/>
            {:else}
                <span class="spinner spinner-primary" style="font-size: 2rem; grid-area: symbol; margin-right: 10px;"/>
            {/if}
            <span class="text">Drop files to Upload or</span> 
            <label for={name} class={`button color-text-inverted button-background-color-primary`}
                class:disabled={working}>
                Browse
            </label>
            {#if accept}
            <small class="accept">Accepts: {formatAccept(accept)}</small>
            {/if}
            <div class="files">
                <!-- displayed the list of files that have been uploaded -->
                {#each files ?? [] as file}
                    <small>
                        <Symbol name="draft" color="primary" variant="outlined" style="font-size: 1rem; padding-right: 0.5em;"/>
                        <span class="file-text">{file.name}</span>
                    </small>
                {/each}
            </div>
        </div>
        
        <!-- input that stores the files, has been made invisible for customisation reasons -->
        <input
            id={name}
            name={name}
            type="file"
            accept={accept}
            bind:files
            {...$$restProps}
            multiple={multiple}
            on:change={change}
            on:dragenter
            on:dragover
            on:dragleave
            on:click
            on:keydown
            on:keyup
            on:keypress
            disabled={working}
            />
    </div>
    <!-- message that shows when there has been an issue with the file upload -->
    {#if failureMessage}
        <span class="color-danger" style="display: block; padding-top: 5px; width: 410px;">{failureMessage}</span>
    {/if}
</div>