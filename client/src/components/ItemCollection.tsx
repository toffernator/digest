import Item from "./Item";


interface ItemCollectionProps {
    items: JSX.Element[],
}

export default function ItemCollection(props: ItemCollectionProps) {
    const itemItems = props.items.map((item) => <li>{item}</li>)

    return (
        <ol className="flex flex-col">{itemItems}</ol>
    );
}