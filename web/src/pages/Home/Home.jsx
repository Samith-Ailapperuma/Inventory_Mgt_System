import React, { useState } from "react";
import { useNavigate } from "react-router-dom"; 
import "./Home.css";

function FeatureCard({ title, description }) {
    return (
        <div>
            <h3>{title}</h3>
            <p>{description}</p>
        </div>
    );
}

export default function Home() {
    const [menuOpen, setMenuOpen] = useState(false);
    const navigate = useNavigate();

    return (
        <div className="home-container">
            <div className="dropdown">
                <div className="hamburger" onClick={() => setMenuOpen(!menuOpen)}>
                    <div className="burger" />
                    <div className="burger" />
                    <div className="burger" />
                </div>

                {menuOpen && (
                    <div className="dropdown-content">
                        <a href="/vendors">Vendors</a>
                        <a href="/items">Items</a>
                    </div>
                )}
            </div>

            <header className="home-header">
                <h1>Inventory Management System</h1>
            </header>

            <section className="features-section">
                <div className="features">
                    <FeatureCard
                        title="Sale"
                        description="Manage sales"
                    />
                    <button onClick={() => navigate("/sales/add")}>Add new sale</button>
                    <button>View Sale Details</button>
                </div>
            </section>
        </div>
    );
}