import { useEffect, useState, useMemo, React } from "react";
import './sales.css';

export default function AddSales() {
    const [items, setItems] = useState([]);
    const [selectedItem, setSelectedItem] = useState("");
    const [quantity, setQuantity] = useState("");
    const [addedSaleItems, setAddedSaleItems] = useState([]);

    const selectedItemData = useMemo(() => {
        return items.find(i => i.Item_Id === selectedItem);
    }, [items, selectedItem]);

    const handleAddItem = async () => {
        if (!selectedItem) {
            alert("Please select an item");
            return;
        }

        if (!quantity || Number(quantity) < 1) {
            alert("Please enter a valid quantity (≥ 1)");
            return;
        }

        if (!selectedItemData) {
            alert("Selected item not found");
            return;
        }

        const newPendingItem = {
            Item_Id: selectedItemData.Item_Id,
            Item_Name: selectedItemData.Item_Name,
            Quantity: Number(quantity),
            Unit_Price: selectedItemData.Unit_Price,
            Total_Price: selectedItemData.Unit_Price * Number(quantity),
        };

        setAddedSaleItems(prev => {
            const similarItemInListIndex = prev.findIndex((item) => item.Item_Id === newPendingItem.Item_Id)
          
            if(similarItemInListIndex === -1) {
                return [...prev, newPendingItem];
            } 

            const newItemsInList = [...addedSaleItems]
            newItemsInList[similarItemInListIndex] = {
                ...newItemsInList[similarItemInListIndex],
                
                Quantity: newItemsInList[similarItemInListIndex].Quantity + newPendingItem.Quantity,
                Total_Price: newItemsInList[similarItemInListIndex].Total_Price + newPendingItem.Total_Price
            }

            return newItemsInList;
        })

        setSelectedItem("");
        setQuantity("");
    }

    const handleCompleteSale = async () => {
        if (addedSaleItems.length === 0) {
            alert("No items added yet");
            return;
        }

        var result = null;
        var addItemRequest = null;
        for (const saleItem of addedSaleItems) {
            try {
                if (result != null) {
                    addItemRequest = {
                        Sale_Id: result.saleId,
                        Qty_Sold: saleItem.Quantity,
                        Item_Id: saleItem.Item_Id
                    }
                } else {
                    addItemRequest = {
                        Qty_Sold: saleItem.Quantity,
                        Item_Id: saleItem.Item_Id
                    }
                }
                console.log(JSON.stringify(addItemRequest));
                const response = await fetch("http://localhost:8080/addItemToSale", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify(addItemRequest),
                });

                if (!response.ok) throw new Error("Failed to save sale");

                result = await response.json();
                console.log("Sale created:", result);

                setAddedSaleItems([]);
                setSelectedItem("");
                setQuantity("");
            } catch (err) {
                console.error(err);
                alert("Failed to complete sale. Check console for details.");
            }
        }

    };

    const handleDeleteItem = (indexToRemove) => {
        setAddedSaleItems(prev =>
            prev.filter((_, index) => index !== indexToRemove)
        );
    };

    const grandTotal = addedSaleItems.reduce(
        (sum, item) => sum + item.Total_Price,
        0
    );

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

            <div className="features">
                <button onClick={handleAddItem}>Add item</button>

                <button
                    onClick={handleCompleteSale}
                    disabled={addedSaleItems.length === 0}
                    style={{ marginTop: "1rem" }}>
                    Complete Sale</button>
            </div>

            {addedSaleItems.length > 0 && (
                <div style={{ marginTop: "1.5rem" }}>
                    <h3>Items to be added:</h3>

                    <table className="sales-table">
                        <thead>
                            <tr>
                                <th></th>
                                <th>Item</th>
                                <th>Quantity</th>
                                <th>Unit Price</th>
                                <th>Total</th>
                            </tr>
                        </thead>

                        <tbody>
                            {addedSaleItems.map((item, index) => (
                                <tr key={index}>
                                    <td>{item.Item_Id}</td>
                                    <td>{item.Item_Name}</td>
                                    <td>{item.Quantity}</td>
                                    <td>${item.Unit_Price}</td>
                                    <td>
                                        ${(item.Quantity * item.Unit_Price).toFixed(2)}
                                    </td>
                                    <td>
                                        <button
                                            onClick={() => handleDeleteItem(index)}
                                            className="delete-btn"
                                        >
                                            Delete
                                        </button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>

                        <tfoot>
                            <tr>
                                <td colSpan="4"></td>
                                <td><strong>${grandTotal.toFixed(2)}</strong></td>
                            </tr>
                        </tfoot>
                    </table>
                </div>
            )}

        </div>
    );

}