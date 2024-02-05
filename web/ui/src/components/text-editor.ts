import Quill, { QuillOptionsStatic } from "quill";
import 'quill/dist/quill.snow.css'

export default class TextEditor extends HTMLElement {
    private static readonly options: QuillOptionsStatic = {
        theme: 'snow',
        modules: {
            toolbar: [
                ['bold', 'italic', 'underline', 'strike'],
                [{ 'list': 'ordered' }, { 'list': 'bullet' }],
                [{ 'align': [] }],
                ['clean'],
            ],
        }
    }

    private instance: Quill | undefined;

    public get quill() { return this.instance!; }

    constructor() {
        super();
    }

    connectedCallback() {
        this.instance = new Quill(this, TextEditor.options);
    }

    adoptedCallback() {
        this.instance?.enable();
    }

    disconnectedCallback() {
        this.instance?.disable();
    }
}
