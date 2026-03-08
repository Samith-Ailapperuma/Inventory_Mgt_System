import { useEffect, useState, useMemo, React } from "react";
import './sales.css';

export default function AddSales() {
    const [items, setItems] = useState([]);
    const [selectedItem, setSelectedItem] = useState("");
    const [quantity, setQuantity] = useState("");
    const selectedItemData = useMemo(() => {
        return items.find(i => i.Item_Id === selectedItem);
    }, [items, selectedItem]);

    useEffect(() => {
        fetch(`http://localhost:8080/getAllItems`)
            .then((res) => res.json())
            .then((data) => setItems(data))
            .catch((err) => console.error(err));
    }, []);

    return (
        <div className="add-sale-item">
            <h1>Add Sale</h1>

            <label className="label">
                Item:
                <select
                    className="select"
                    value={selectedItem}
                    onChange={(e) => setSelectedItem(e.target.value)}>
                    <option value="" disabled>
                        Select an item
                    </option>
                    {items.map((item) => (
                        <option key={item.Item_Id} value={item.Item_Id}>{item.Item_Name}</option>
                    ))}
                </select>
            </label>

            <label class="label">
                Amount/No.of Units:
                <input
                    type="number"
                    min="1"
                    className="select"
                    value={quantity}
                    onChange={(e) => setQuantity(e.target.value)}
                    placeholder="Enter quantity"
                />
            </label>

            <label class="label">
                Price:
                <div className="select">{selectedItemData?.Unit_Price}
                </div>
            </label>

            <button>Add item</button>

            <button>Complete Sale</button>
        </div>
    );

}