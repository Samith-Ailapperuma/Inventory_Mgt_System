import React from "react";
import "./Home.css";

// interface FeatureCardProps {
//     title: string;
//     description: string
// }

function FeatureCard({title, description}) {
    return (
        <div>
            <h3>{title}</h3>
            <p>{description}</p>
        </div>
    );
}

export default function Home() {
    return (
        <div className="home-container">
            <header className="home-header">
                <h1>Inventory Management System</h1>
            </header>

            <section className="features-section">
                <h2>Features</h2>
                <div className="features">
                    <FeatureCard
                        title="Vendors"
                        description="Handle all activities related to vendors"
                    />
                    <FeatureCard
                        title="Items"
                        description="Manage items"
                    />
                    <FeatureCard
                        title="Sale"
                        description="Manage sales"
                    />
                </div>
            </section>
        </div>
    );
}