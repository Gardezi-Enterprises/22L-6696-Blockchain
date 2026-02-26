import React, { useState, useEffect } from 'react';
import AddTransaction from './components/AddTransaction';
import MineBlock from './components/MineBlock';
import ViewBlockchain from './components/ViewBlockchain';
import SearchBlock from './components/SearchBlock';
import PendingTransactions from './components/PendingTransactions';
import './App.css';

const API = 'http://localhost:8080';

export default function App() {
  const [activeTab, setActiveTab] = useState('view');
  const [chain, setChain] = useState([]);
  const [pending, setPending] = useState([]);
  const [chainLength, setChainLength] = useState(0);
  const [isValid, setIsValid] = useState(true);
  const [loading, setLoading] = useState(false);
  const [notification, setNotification] = useState(null);

  const showNotification = (msg, type = 'success') => {
    setNotification({ msg, type });
    setTimeout(() => setNotification(null), 4000);
  };

  const fetchChain = async () => {
    setLoading(true);
    try {
      const res = await fetch(`${API}/chain`);
      const data = await res.json();
      setChain(data.chain || []);
      setChainLength(data.length || 0);
      setIsValid(data.valid);
    } catch {
      showNotification('Cannot connect to Go server. Make sure it runs on :8080', 'error');
    } finally {
      setLoading(false);
    }
  };

  const fetchPending = async () => {
    try {
      const res = await fetch(`${API}/pending`);
      const data = await res.json();
      setPending(data.pendingTransactions || []);
    } catch {}
  };

  useEffect(() => {
    fetchChain();
    fetchPending();
  }, []);

  const refresh = () => {
    fetchChain();
    fetchPending();
  };

  const tabs = [
    { id: 'view',    label: '🔗 View Chain' },
    { id: 'add',     label: '➕ Add Transaction' },
    { id: 'pending', label: `⏳ Pending (${pending.length})` },
    { id: 'mine',    label: '⛏ Mine Block' },
    { id: 'search',  label: '🔍 Search' },
  ];

  return (
    <div className="app">
      {/* Header */}
      <header className="header">
        <div className="header-inner">
          <div className="logo">
            <span className="logo-icon">⛓</span>
            <div>
              <h1>Blockchain Explorer</h1>
              <p>Go + React · Proof of Work · Merkle Tree</p>
            </div>
          </div>
          <div className="stats">
            <div className="stat">
              <span className="stat-value">{chainLength}</span>
              <span className="stat-label">Blocks</span>
            </div>
            <div className="stat">
              <span className="stat-value">{pending.length}</span>
              <span className="stat-label">Pending Tx</span>
            </div>
            <div className={`stat validity ${isValid ? 'valid' : 'invalid'}`}>
              <span className="stat-value">{isValid ? '✓' : '✗'}</span>
              <span className="stat-label">{isValid ? 'Valid' : 'Tampered'}</span>
            </div>
            <button className="refresh-btn" onClick={refresh} disabled={loading}>
              {loading ? '...' : '↻ Refresh'}
            </button>
          </div>
        </div>
      </header>

      {/* Notification */}
      {notification && (
        <div className={`notification ${notification.type}`}>
          {notification.msg}
        </div>
      )}

      {/* Tabs */}
      <nav className="tabs">
        {tabs.map(t => (
          <button
            key={t.id}
            className={`tab ${activeTab === t.id ? 'active' : ''}`}
            onClick={() => setActiveTab(t.id)}
          >
            {t.label}
          </button>
        ))}
      </nav>

      {/* Content */}
      <main className="content">
        {activeTab === 'view' && (
          <ViewBlockchain chain={chain} loading={loading} />
        )}
        {activeTab === 'add' && (
          <AddTransaction
            api={API}
            onSuccess={(msg) => {
              showNotification(msg);
              fetchPending();
            }}
            onError={(msg) => showNotification(msg, 'error')}
          />
        )}
        {activeTab === 'pending' && (
          <PendingTransactions pending={pending} onRefresh={fetchPending} />
        )}
        {activeTab === 'mine' && (
          <MineBlock
            api={API}
            pendingCount={pending.length}
            onSuccess={(msg) => {
              showNotification(msg);
              refresh();
            }}
            onError={(msg) => showNotification(msg, 'error')}
          />
        )}
        {activeTab === 'search' && (
          <SearchBlock api={API} onError={(msg) => showNotification(msg, 'error')} />
        )}
      </main>
    </div>
  );
}
