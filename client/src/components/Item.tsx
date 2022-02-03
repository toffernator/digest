import Tag from "./Tag";

interface ItemProps {
    title: string;
    src: string;
    tags: string[];
    url: string;
}

export default function Item(props: ItemProps) {
    let tagItems = props.tags.map((tag) => <li><Tag name={tag} /></li>)
    tagItems = [<li><Tag name={props.src} /></li>, ...tagItems]

    return (
        <div className="flex-1 m-1 p-1 rounded-md bg-gray-100 shadow-md">
            <a href={props.url} className="hover:underline hover:text-emerald-500"><h2 className="text-xl font-serif truncate">{props.title}</h2></a>
            <ol className="flex">{tagItems}</ol>
        </div>
    );
}