import React from "react";

interface TagProps {
    name: string;
}

export default function Tag(props: TagProps) {
        return (
            <div className="mx-0.5 p-1 w-24 rounded-3xl text-center bg-gray-300">
                <p className="select-none truncate">{props.name}</p>
            </div>
        );
}
