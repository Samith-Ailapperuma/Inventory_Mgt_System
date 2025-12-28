import { useEffect, useState, React } from "react";

export default function AddSales() {
    const [items, setItems] = useState([]);
    const [selectedItem, setSelectedItem] = useState("");

    useEffect(() => {
        fetch(`http://localhost:8080/getAllItems`)
            .then((res) => res.json())
            .then((data) => setItems(data))
            .catch((err) => console.error(err));
    }, []);

    return (
        <div>
            <h1>Add Sale</h1>

            <label>
                Item:
                <select 
                    value={selectedItem}
                    onChange={(e) => setSelectedItem(e.target.value)}>
                    {items.map((item) => (
                        <option key={item.Item_Id} value={item.Item_Id}>{item.Item_Name}</option>
                    ))}
                </select>
            </label>
        </div>
    );

}