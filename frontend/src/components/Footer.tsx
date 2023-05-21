import React from 'react';

const Footer: React.FC = () => {
    return (
        <footer className="bg-gray-800 py-4">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <p className="text-center text-gray-300 text-sm">
                    &copy; {new Date().getFullYear()} Your Website Name. All rights reserved.
                </p>
            </div>
        </footer>
    );
};

export default Footer;
