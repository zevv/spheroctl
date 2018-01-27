#!/usr/bin/python3

import sys, pygame
from math import *
from sphero_sprk import Sphero	

    
orb = Sphero("C2:A2:3D:46:C2:87")
orb.connect()
orb.set_rgb_led(10,0,0)
orb.command("21", [255])

pygame.init()

size = width, height = 800, 600
speed = [2, 2]
black = 0, 0, 0

screen = pygame.display.set_mode(size)

pygame.time.set_timer(pygame.USEREVENT, 100)
    
ang = 0
dang = 0
r = 0
down = False

while 1:

    screen.fill(black)
    pygame.draw.circle(screen, (100, 100, 100), (400, 300),  20, 5)
    pygame.draw.circle(screen, (100, 100, 100), (400, 300), 250, 5)

    for event in pygame.event.get():

        if event.type == pygame.KEYDOWN:
            if event.key == ord('q'):
                exit(0)
            if event.key == 27:
                exit(0)
            if event.key == ord('f'):
                pygame.display.toggle_fullscreen()
            if event.key == pygame.K_LEFT:
                ang = ang - 30
            if event.key == pygame.K_RIGHT:
                ang = ang + 30
            if event.key == pygame.K_UP:
                r = r + 10
            if event.key == pygame.K_DOWN:
                r = r - 10
            if event.key == 32:
                r = 0
            if event.key == ord(','):
                dang = dang - 10
            if event.key == ord('.'):
                dang = dang + 10
            if event.key == ord('a'):
                ang = -90
                r = 250
            if event.key == ord('w'):
                ang = 0
                r = 250
            if event.key == ord('d'):
                ang = 90
                r = 250
            if event.key == ord('x'):
                ang = 180
                r = 250
       
        if event.type == pygame.KEYUP:
            if event.key == ord('a'):
                r = 0
            if event.key == ord('w'):
                r = 0
            if event.key == ord('d'):
                r = 0
            if event.key == ord('x'):
                r = 0


        if event.type == pygame.USEREVENT:
            rn = r / 256
            speed = int(255 * (rn * rn * rn))
            orb.roll(speed, (ang + dang) % 360)

        if event.type == pygame.QUIT:
            orb.roll(0, 0)
            sys.exit()
        
        if event.type == pygame.MOUSEBUTTONDOWN:
            down = True
        
        if event.type == pygame.MOUSEBUTTONUP:
            r = 0
            down = False

        if event.type == pygame.MOUSEMOTION:
            (x, y) = event.pos
            a = atan2(y-300, x-400)
            r = int(hypot(y-300, x-400))
            ang = int(a / 3.1415 * 180 + 90) % 360

            if r > 250:
                r = 250

            if not down:
                r = 0

    a = (ang - 90) / 180 * 3.1415
    x = int(400 + cos(a) * r)
    y = int(300 + sin(a) * r)
            
    pygame.draw.lines(screen, (0, 0, 255), False, [(400, 300), (x, y)], 10)
    pygame.draw.circle(screen, (255, 0, 0), (x, y), 20, 5)
    pygame.display.update()

    pygame.display.flip()
    pygame.time.wait(20)

# End

