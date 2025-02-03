import React from 'react';
import { render, screen, fireEvent } from '@testing-library/react';
import {describe,it,expect} from "vitest";
import '@testing-library/jest-dom/vitest';
import App from './App';

describe('App Component', () => {
  it('renders the logos, header, and initial button', () => {
    render(<App />);

    expect(screen.getByAltText(/Vite logo/i)).toBeInTheDocument();
    expect(screen.getByAltText(/React logo/i)).toBeInTheDocument();
    expect(screen.getByText(/Vite \+ React/i)).toBeInTheDocument();
    expect(screen.getByRole('button', { name: /count is 0/i })).toBeInTheDocument();
  });
});
