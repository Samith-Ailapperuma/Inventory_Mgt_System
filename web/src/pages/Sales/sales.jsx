import React from "react";

export default function Sales() {
    return (
        <div>
            <h1>Add Sale</h1>

            <label>
                Item:
                <select name="selectedItem">
                    <option value="item 1">Item 1</option>
                    <option value="item 2">Item 2</option>
                </select>
            </label>
        </div>
    );

}