import React, { useState } from 'react';
import { generatePasswords } from './api';
import { Settings, Copy, RefreshCw, ShieldAlert, ShieldCheck, Check } from 'lucide-react';
import { clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';

function cn(...inputs) {
  return twMerge(clsx(inputs));
}

const LANGUAGES = [
  { value: 'en', label: 'English' },
  { value: 'fi', label: 'Finnish' },
  { value: 'fr', label: 'French' },
];

function App() {
  const [mode, setMode] = useState('password'); // password | passphrase
  const [loading, setLoading] = useState(false);
  const [results, setResults] = useState([]);
  const [warnings, setWarnings] = useState([]);

  // Password Config
  const [length, setLength] = useState(16);
  const [count, setCount] = useState(1);
  const [useUpper, setUseUpper] = useState(true);
  const [useDigits, setUseDigits] = useState(true);
  const [useSpecial, setUseSpecial] = useState(true);
  const [excludeHomo, setExcludeHomo] = useState(false);

  // Passphrase Config
  const [wordCount, setWordCount] = useState(4);
  const [separator, setSeparator] = useState('-');
  const [capitalize, setCapitalize] = useState(true);
  const [includeNumber, setIncludeNumber] = useState(true);
  const [language, setLanguage] = useState('en');

  const handleGenerate = async () => {
    setLoading(true);
    setResults([]);
    setWarnings([]);

    const payload = {
      type: mode,
      count: parseInt(count),
    };

    if (mode === 'password') {
      payload.length = parseInt(length);
      payload.include_uppercase = useUpper;
      payload.include_digits = useDigits;
      payload.include_special = useSpecial;
      payload.exclude_homoglyphs = excludeHomo;
    } else {
      payload.word_count = parseInt(wordCount);
      payload.separator = separator;
      payload.capitalize = capitalize;
      payload.include_number = includeNumber;
      payload.language = language;
    }

    try {
      const data = await generatePasswords(payload);
      setResults(data.passwords || []);
      setWarnings(data.warnings || []);
    } catch (err) {
      console.error(err);
      setWarnings([err.message]);
    } finally {
      setLoading(false);
    }
  };

  const copyToClipboard = (text) => {
    navigator.clipboard.writeText(text);
    // Could add toast here
  };

  return (
    <div className="min-h-screen bg-gray-950 text-gray-100 flex flex-col items-center py-10 px-4">
      <header className="mb-8 text-center">
        <h1 className="text-4xl font-bold bg-gradient-to-r from-indigo-400 to-emerald-400 bg-clip-text text-transparent mb-2">
          PWGen Go
        </h1>
        <p className="text-gray-400">Secure Password & Passphrase Generator</p>
      </header>

      <div className="w-full max-w-md bg-gray-900 rounded-2xl shadow-2xl overflow-hidden border border-gray-800">

        {/* Tabs */}
        <div className="flex border-b border-gray-800">
          <button
            onClick={() => setMode('password')}
            className={cn(
              "flex-1 py-4 text-sm font-medium transition-colors",
              mode === 'password' ? "bg-gray-800 text-white" : "text-gray-400 hover:text-gray-200 hover:bg-gray-800/50"
            )}
          >
            Password
          </button>
          <button
            onClick={() => setMode('passphrase')}
            className={cn(
              "flex-1 py-4 text-sm font-medium transition-colors",
              mode === 'passphrase' ? "bg-gray-800 text-white" : "text-gray-400 hover:text-gray-200 hover:bg-gray-800/50"
            )}
          >
            Passphrase
          </button>
        </div>

        {/* Config Area */}
        <div className="p-6 space-y-6">

          {/* Common: Count */}
          <div className="space-y-2">
            <div className="flex justify-between text-sm">
              <label className="text-gray-400">Count</label>
              <span className="font-mono">{count}</span>
            </div>
            <input
              type="range"
              min="1"
              max="5"
              value={count}
              onChange={(e) => setCount(e.target.value)}
              className="range-slider"
            />
          </div>

          {mode === 'password' ? (
            <>
              {/* Length */}
              <div className="space-y-2">
                <div className="flex justify-between text-sm">
                  <label className="text-gray-400">Length</label>
                  <span className="font-mono">{length}</span>
                </div>
                <input
                  type="range"
                  min="8"
                  max="64"
                  value={length}
                  onChange={(e) => setLength(e.target.value)}
                  className="range-slider"
                />
              </div>

              {/* Toggles */}
              <div className="grid grid-cols-2 gap-4">
                <Toggle label="Uppercase (A-Z)" checked={useUpper} onChange={setUseUpper} />
                <Toggle label="Digits (0-9)" checked={useDigits} onChange={setUseDigits} />
                <Toggle label="Special (!@#)" checked={useSpecial} onChange={setUseSpecial} />
                <Toggle label="No Homoglyphs" checked={excludeHomo} onChange={setExcludeHomo} />
              </div>
            </>
          ) : (
            <>
              {/* Word Count */}
              <div className="space-y-2">
                <div className="flex justify-between text-sm">
                  <label className="text-gray-400">Words</label>
                  <span className="font-mono">{wordCount}</span>
                </div>
                <input
                  type="range"
                  min="3"
                  max="10"
                  value={wordCount}
                  onChange={(e) => setWordCount(e.target.value)}
                  className="range-slider"
                />
              </div>

              {/* Separator */}
              <div className="space-y-2">
                <label className="text-gray-400 text-sm">Separator</label>
                <input
                  type="text"
                  maxLength="1"
                  value={separator}
                  onChange={(e) => setSeparator(e.target.value)}
                  className="w-full bg-gray-800 border-none rounded py-2 px-3 text-center font-mono focus:ring-2 focus:ring-indigo-500 outline-none"
                />
              </div>

              {/* Language */}
              <div className="space-y-2">
                <label className="text-gray-400 text-sm">Language</label>
                <select
                  value={language}
                  onChange={(e) => setLanguage(e.target.value)}
                  className="w-full bg-gray-800 border-none rounded py-2 px-3 text-white focus:ring-2 focus:ring-indigo-500 outline-none"
                >
                  {LANGUAGES.map(l => <option key={l.value} value={l.value}>{l.label}</option>)}
                </select>
              </div>

              {/* Toggles */}
              <div className="grid grid-cols-2 gap-4">
                <Toggle label="Capitalize" checked={capitalize} onChange={setCapitalize} />
                <Toggle label="Include Number" checked={includeNumber} onChange={setIncludeNumber} />
              </div>
            </>
          )}

          {/* Action */}
          <button
            onClick={handleGenerate}
            disabled={loading}
            className="w-full py-3 bg-indigo-600 hover:bg-indigo-500 active:bg-indigo-700 rounded-xl font-bold transition-all shadow-lg shadow-indigo-500/20 disabled:opacity-50 flex items-center justify-center gap-2"
          >
            {loading ? <RefreshCw className="animate-spin w-5 h-5" /> : <Settings className="w-5 h-5" />}
            Generate
          </button>

        </div>
      </div>

      {/* Output Area */}
      {results.length > 0 && (
        <div className="w-full max-w-md mt-8 space-y-4 animate-fade-in-up">
          {warnings.length > 0 && (
            <div className="bg-yellow-900/30 text-yellow-200 p-4 rounded-xl border border-yellow-700/50 flex gap-3 text-sm">
              <ShieldAlert className="shrink-0 w-5 h-5" />
              <ul>
                {warnings.map((w, i) => <li key={i}>{w}</li>)}
              </ul>
            </div>
          )}

          {results.map((res, idx) => (
            <div key={idx} className="group relative bg-gray-900 p-4 rounded-xl border border-gray-800 flex items-center justify-between hover:border-indigo-500/50 transition-colors">
              <div className="font-mono text-lg break-all mr-4 text-emerald-400">
                {res}
              </div>
              <div className="flex gap-2">
                <button
                  onClick={() => copyToClipboard(res)}
                  className="p-2 text-gray-500 hover:text-white rounded-lg hover:bg-gray-800 transition-colors"
                >
                  <Copy className="w-5 h-5" />
                </button>
              </div>
            </div>
          ))}
        </div>
      )}

    </div>
  );
}

function Toggle({ label, checked, onChange }) {
  return (
    <button
      onClick={() => onChange(!checked)}
      className={cn(
        "flex flex-col items-center justify-center py-3 px-2 rounded-lg border transition-all",
        checked
          ? "bg-indigo-600/10 border-indigo-500/50 text-indigo-300"
          : "bg-gray-800 border-transparent text-gray-500 hover:bg-gray-750"
      )}
    >
      <span className="text-xs font-medium mb-1">{label}</span>
      {checked ? <Check className="w-4 h-4" /> : <div className="w-4 h-4 rounded-full border border-gray-600" />}
    </button>
  );
}

export default App;
