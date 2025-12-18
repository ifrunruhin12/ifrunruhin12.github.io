// Page Flip Animation System - FIXED VERSION
import { pageState, updatePageState } from "./pageState.js";
import { audioSystem } from "./audioSystem.js";

export function performPageFlip(targetPageId, direction) {
  if (pageState.isAnimating) return;

  pageState.isAnimating = true;
  pageState.animationDirection = direction;

  const currentPageElement = document.getElementById(
    pageState.currentPage + "-page",
  );
  const targetPageElement = document.getElementById(targetPageId + "-page");
  const mainContainer = document.querySelector(".newspaper-main");

  if (!currentPageElement || !targetPageElement) {
    pageState.isAnimating = false;
    return;
  }

  // Add direction class to container
  mainContainer.classList.add(`page-flip-${direction}`);

  // Perform the flip
  perform3DPageFlip(
    currentPageElement,
    targetPageElement,
    targetPageId,
    direction,
  );
}

function perform3DPageFlip(currentPage, targetPage, targetPageId, direction) {
  console.log(
    "Starting page flip:",
    direction,
    "from",
    currentPage.id,
    "to",
    targetPage.id,
  );

  // Play sound
  audioSystem.playPageFlip();

  // Remove any existing animation classes
  currentPage.classList.remove("flipping-out", "flipping-in");
  targetPage.classList.remove("flipping-out", "flipping-in");

  // Force reflow
  void currentPage.offsetHeight;
  void targetPage.offsetHeight;

  // Make target page visible but transparent
  targetPage.style.visibility = "visible";
  targetPage.style.opacity = "0";
  targetPage.style.zIndex = "10";

  // Set initial transform for target page based on direction
  if (direction === "backward") {
    targetPage.style.transform = "perspective(2000px) rotateY(-180deg)";
  } else {
    targetPage.style.transform = "perspective(2000px) rotateY(180deg)";
  }

  // Start animations on next frame
  requestAnimationFrame(() => {
    currentPage.classList.add("flipping-out");
    targetPage.classList.add("flipping-in");
    currentPage.style.zIndex = "15";
  });

  // Animation duration matches CSS (1000ms)
  const animationDuration = 1000;

  setTimeout(() => {
    // Clean up current page
    currentPage.classList.remove("active", "flipping-out");
    currentPage.style.visibility = "hidden";
    currentPage.style.opacity = "1";
    currentPage.style.transform = "";
    currentPage.style.zIndex = "";

    // Finalize target page
    targetPage.classList.remove("flipping-in");
    targetPage.classList.add("active");
    targetPage.style.visibility = "visible";
    targetPage.style.opacity = "1";
    targetPage.style.transform = "";
    targetPage.style.zIndex = "";

    // Clean up container
    const mainContainer = document.querySelector(".newspaper-main");
    mainContainer.classList.remove(`page-flip-${direction}`);

    // Update state
    updatePageState(targetPageId);
    smoothScrollToTop();
    pageState.isAnimating = false;

    console.log("Animation completed. Active page:", targetPageId);
  }, animationDuration);
}

function smoothScrollToTop() {
  const startPosition = window.pageYOffset;
  const startTime = performance.now();
  const duration = 400;

  function easeInOutCubic(t) {
    return t < 0.5 ? 4 * t * t * t : (t - 1) * (2 * t - 2) * (2 * t - 2) + 1;
  }

  function scrollAnimation(currentTime) {
    const elapsed = currentTime - startTime;
    const progress = Math.min(elapsed / duration, 1);
    const easeProgress = easeInOutCubic(progress);

    window.scrollTo(0, startPosition * (1 - easeProgress));

    if (progress < 1) {
      requestAnimationFrame(scrollAnimation);
    }
  }

  requestAnimationFrame(scrollAnimation);
}
