import { render, screen } from "@testing-library/preact";
import { describe, expect, it } from "vitest";
import { Spinner } from "./index";

describe("Spinner", () => {
  it("renders when isOpen is true", () => {
    render(<Spinner isOpen={true} />);
    const spinner = screen.getByRole("progressbar");
    expect(spinner).toBeInTheDocument();
  });

  it("does not render when isOpen is false", () => {
    render(<Spinner isOpen={false} />);
    const spinner = screen.queryByRole("progressbar");
    expect(spinner).not.toBeInTheDocument();
  });
});
