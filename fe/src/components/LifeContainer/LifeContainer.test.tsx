import { render, screen } from "@testing-library/preact";
import { describe, expect, it } from "vitest";
import { LifeContainer } from "./index";

describe("LifeContainer", () => {
  it("renders no hearts when lives is 0", () => {
    const { container } = render(<LifeContainer lives={0} />);
    const hearts = container.querySelectorAll('span[aria-hidden="true"]');
    expect(hearts.length).toBe(0);
  });

  it("renders a single heart when lives is 1", () => {
    const { container } = render(<LifeContainer lives={1} />);
    const hearts = container.querySelectorAll('span[aria-hidden="true"]');
    expect(hearts.length).toBe(1);
    expect(hearts[0].textContent).toBe("♥");
  });

  it("renders multiple hearts when lives is greater than 1", () => {
    const { container } = render(<LifeContainer lives={3} />);
    const hearts = container.querySelectorAll('span[aria-hidden="true"]');
    expect(hearts.length).toBe(3);
  });

  it("renders correct number of hearts for various life counts", () => {
    const testCases = [5, 10, 7];

    testCases.forEach((lives) => {
      const { container, unmount } = render(<LifeContainer lives={lives} />);
      const hearts = container.querySelectorAll('span[aria-hidden="true"]');
      expect(hearts.length).toBe(lives);
      unmount();
    });
  });

  it("has proper accessibility attributes", () => {
    render(<LifeContainer lives={3} />);
    const status = screen.getByRole("status");
    expect(status).not.toBeNull();
    expect(status.getAttribute("aria-label")).toBe("3 lives remaining");
  });

  it("updates aria-label correctly for different life counts", () => {
    const { rerender } = render(<LifeContainer lives={5} />);
    let status = screen.getByRole("status");
    expect(status.getAttribute("aria-label")).toBe("5 lives remaining");

    rerender(<LifeContainer lives={1} />);
    status = screen.getByRole("status");
    expect(status.getAttribute("aria-label")).toBe("1 lives remaining");
  });
});
