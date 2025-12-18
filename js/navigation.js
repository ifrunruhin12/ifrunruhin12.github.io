// Navigation System
import { pageState, getCurrentPageOrder, getPageOrder } from './pageState.js';
import { performPageFlip } from './animations.js';

export function initializeNavigation() {
    const navLinks = document.querySelectorAll('.nav-link');
    
    navLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            
            if (pageState.isAnimating) return;
            
            const targetPage = this.getAttribute('data-page');
            if (targetPage === pageState.currentPage) return;
            
            const currentPageOrder = getCurrentPageOrder();
            const targetPageOrder = getPageOrder(targetPage);
            const direction = targetPageOrder > currentPageOrder ? 'forward' : 'backward';
            
            performPageFlip(targetPage, direction);
            updateActiveNavLink(this);
            
            if (window.innerWidth <= 768) {
                closeMobileMenu();
            }
        });
    });
    
    // Initialize with home page active
    const homePage = document.getElementById('home-page');
    if (homePage) {
        homePage.classList.add('active');
        homePage.style.visibility = 'visible';
        homePage.style.opacity = '1';
    }
}

export function initializeMobileMenu() {
    const navToggle = document.querySelector('.nav-toggle');
    const newspaperNav = document.querySelector('.newspaper-nav');
    
    if (navToggle && newspaperNav) {
        navToggle.addEventListener('click', toggleMobileMenu);
        
        document.addEventListener('click', function(e) {
            if (window.innerWidth <= 768 && 
                !e.target.closest('.newspaper-nav') && 
                !e.target.closest('.nav-toggle') &&
                newspaperNav.classList.contains('active')) {
                closeMobileMenu();
            }
        });
        
        window.addEventListener('resize', function() {
            if (window.innerWidth > 768) {
                closeMobileMenu();
            }
        });
    }
}

export function initializeKeyboardNavigation() {
    document.addEventListener('keydown', function(e) {
        if (e.key === 'Escape') {
            closeMobileMenu();
            return;
        }
        
        if (pageState.isAnimating) return;
        
        if (e.key === 'ArrowLeft' || e.key === 'ArrowRight') {
            e.preventDefault();
            
            const currentOrder = getCurrentPageOrder();
            let targetOrder;
            
            if (e.key === 'ArrowLeft' && currentOrder > 0) {
                targetOrder = currentOrder - 1;
            } else if (e.key === 'ArrowRight' && currentOrder < pageState.pages.length - 1) {
                targetOrder = currentOrder + 1;
            }
            
            if (targetOrder !== undefined) {
                const targetPage = pageState.pages.find(p => p.order === targetOrder);
                if (targetPage) {
                    const targetNavLink = document.querySelector(`[data-page="${targetPage.id}"]`);
                    if (targetNavLink) {
                        targetNavLink.click();
                    }
                }
            }
        }
    });
}

function updateActiveNavLink(activeLink) {
    const navLinks = document.querySelectorAll('.nav-link');
    navLinks.forEach(nav => nav.classList.remove('active'));
    activeLink.classList.add('active');
}

function toggleMobileMenu() {
    const newspaperNav = document.querySelector('.newspaper-nav');
    newspaperNav.classList.toggle('active');
    
    if (newspaperNav.classList.contains('active')) {
        document.body.style.overflow = 'hidden';
    } else {
        document.body.style.overflow = '';
    }
}

export function closeMobileMenu() {
    const newspaperNav = document.querySelector('.newspaper-nav');
    if (newspaperNav) {
        newspaperNav.classList.remove('active');
        document.body.style.overflow = '';
    }
}