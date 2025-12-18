// Page Flip Animation System
import { pageState, updatePageState } from './pageState.js';
import { audioSystem } from './audioSystem.js';

export function performPageFlip(targetPageId, direction) {
    if (pageState.isAnimating) return;
    
    pageState.isAnimating = true;
    pageState.animationDirection = direction;
    
    const currentPageElement = document.getElementById(pageState.currentPage + '-page');
    const targetPageElement = document.getElementById(targetPageId + '-page');
    const mainContainer = document.querySelector('.newspaper-main');
    
    if (!currentPageElement || !targetPageElement) {
        pageState.isAnimating = false;
        return;
    }
    
    mainContainer.classList.add(`page-flip-${direction}`);
    
    // Force 3D animations for debugging
    perform3DPageFlip(currentPageElement, targetPageElement, targetPageId, direction);
}

function perform3DPageFlip(currentPage, targetPage, targetPageId, direction) {
    console.log('Starting 3D page flip animation', direction);
    console.log('Current page:', currentPage.id, 'Target page:', targetPage.id);
    audioSystem.playPageFlip();
    
    // Force a reflow to ensure classes are applied
    currentPage.offsetHeight;
    targetPage.offsetHeight;
    
    // Add animation classes
    currentPage.classList.add('flipping-out');
    targetPage.classList.add('flipping-in');
    targetPage.style.visibility = 'visible';
    
    // Ensure target page starts with correct initial state for animation
    if (direction === 'backward') {
        targetPage.style.transform = 'rotateY(-180deg)';
    } else {
        targetPage.style.transform = 'rotateY(180deg)';
    }
    targetPage.style.opacity = '0';
    
    console.log('Classes added:', currentPage.classList.contains('flipping-out'), targetPage.classList.contains('flipping-in'));
    
    // Add the paper bend effect
    currentPage.style.animation = 'paperBend 1.2s ease-in-out';
    
    const animationDuration = 1200;
    
    setTimeout(() => {
        // Clean up current page
        currentPage.classList.remove('active', 'flipping-out', 'test-animation');
        currentPage.style.visibility = 'hidden';
        currentPage.style.animation = '';
        
        // Show target page
        targetPage.classList.remove('flipping-in');
        targetPage.classList.add('active');
        targetPage.style.visibility = 'visible';
        targetPage.style.opacity = '1';
        targetPage.style.transform = ''; // Clear any inline transforms
        
        // Clean up container
        const mainContainer = document.querySelector('.newspaper-main');
        mainContainer.classList.remove(`page-flip-${direction}`);
        
        updatePageState(targetPageId);
        smoothScrollToTop();
        pageState.isAnimating = false;
        
        console.log('Animation completed, target page should be visible:', targetPageId);
    }, animationDuration);
}

function performSimpleTransition(currentPage, targetPage, targetPageId) {
    currentPage.style.opacity = '0';
    
    setTimeout(() => {
        currentPage.classList.remove('active');
        currentPage.style.visibility = 'hidden';
        currentPage.style.opacity = '1';
        
        targetPage.classList.add('active');
        targetPage.style.visibility = 'visible';
        targetPage.style.opacity = '1';
        
        const mainContainer = document.querySelector('.newspaper-main');
        mainContainer.classList.remove(`page-flip-${pageState.animationDirection}`);
        
        updatePageState(targetPageId);
        
        window.scrollTo({
            top: 0,
            behavior: 'smooth'
        });
        
        pageState.isAnimating = false;
    }, 300);
}

function smoothScrollToTop() {
    const startPosition = window.pageYOffset;
    const startTime = performance.now();
    const duration = 600;
    
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

function checkTransformSupport() {
    const testElement = document.createElement('div');
    const transforms = [
        'transform',
        'WebkitTransform',
        'MozTransform',
        'msTransform',
        'OTransform'
    ];
    
    for (let i = 0; i < transforms.length; i++) {
        if (testElement.style[transforms[i]] !== undefined) {
            return true;
        }
    }
    return false;
}