import { React, useState } from 'react'
import './sales.css';

export default function ViewSales() {

    const [sales, setSales] = useState([]);
    const [loading, setLoading] = useState(false); 
    const [error, setError] = useState('');
    const [viewSalesClicked, setViewSalesClicked] = useState(false);

    const handleViewSales = async () => {
        try {
            const response = await fetch('http://localhost:8080/allSales', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }

            const data = await response.json();
            setSales(data);
            setViewSalesClicked(true);
        } catch (err) {
            console.err('Error fetching sales:', err);
            setError('Failed to load sales data.');
        } finally { 
            setLoading(false); 
        }
    }

    return (
        <div className="add-sale-item">
            <h1>View Sale Details</h1>

            <div>
                <button onClick={handleViewSales}>View All Sales</button>
            </div>

            {loading && 
                <p>Loading sales...</p>} 
            {error && 
                <p>{error}</p>}

            {sales.length > 0 && (
                <table className="view-sales-table"> 
                    <thead> 
                        <tr> 
                            <th>Sale ID</th> 
                            <th>Amount</th> 
                            <th>Test</th> 
                        </tr> 
                    </thead> 
                    
                    <tbody> 
                        {sales.map((sale) => (
                            <tr>
                                <td>{sale.Sale_Id}</td>
                                <td>{sale.Sale_Amount}</td>
                                <td><button>View Sale Details</button></td> 
                            </tr>))} 
                    </tbody> 
                </table>
            )}

            {!loading && sales.length === 0 && !error && viewSalesClicked && ( <p>No sales to display.</p> )}

        </div>
    )
}