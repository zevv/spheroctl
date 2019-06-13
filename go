#!/usr/bin/python3

import sys, pygame
from math import *
from sphero_sprk import Sphero	

    
orb = Sphero("C2:A2:3D:46:C2:87")
orb.connect()
orb.set_rgb_led(10,0,0)
orb.command("21", [100])

pygame.init()

size = width, height = 800, 600
speed = [2, 2]
black = 0, 0, 0

screen = pygame.display.set_mode(size)

pygame.time.set_timer(pygame.USEREVENT, 100)
    
ang = 0
ang2 = 10
dang = 0
r = 0
down = False
stab = True

acc = []
gyro = []
imu = []

def on_accel(v):
    global acc
    acc.append(v)
    if len(acc) > 80:
        del acc[0]

def on_gyro(v):
    global gyro
    gyro.append(v)
    if len(gyro) > 80:
        del gyro[0]

def on_imu(v):
    global ang2
    ang2 = -v['yaw']
    imu.append(v)
    if len(imu) > 80:
        del imu[0]

orb.start_accel_callback(20, on_accel)
orb.start_gyro_callback(20, on_gyro)
orb.start_IMU_callback(20, on_imu)


def graph(gs, k, y, scale=1):
    if len(gs) > 1:
        ls = []
        x = 0
        for g in gs:
            ls.append((x, y + g[k] * scale))
            x = x + 10
        pygame.draw.lines(screen, (100, 100, 100), False, ls, 1)


def trick1():
    orb.roll(0, 0)
    pygame.time.wait(300)
    for i in range(0, 36*4):
        orb.roll(255, (i*10) % 360)
        pygame.time.wait(30)
    orb.roll(0, 0)

while 1:

    screen.fill(black)
    pygame.draw.circle(screen, (100, 100, 100), (400, 300),  20, 5)
    pygame.draw.circle(screen, (100, 100, 100), (400, 300), 250, 5)
    
    graph(acc, "x",  60, 0.01)
    graph(acc, "y",  80, 0.01)
    graph(acc, "z", 100, 0.01)
    
    graph(imu, "pitch", 260)
    graph(imu, "roll", 280)
    graph(imu, "yaw", 300)

#    graph(gyro, 200)
#    graph(imu, 300)

    for event in pygame.event.get():

        if event.type == pygame.KEYDOWN:
            if event.key == ord('s'):
                stab = False
                r = 0
            if event.key == ord('q'):
                exit(0)
            if event.key == 27:
                exit(0)
            if event.key == ord('f'):
                pygame.display.toggle_fullscreen()
            if event.key == 32:
                stab = True
                r = 0
            if event.key == pygame.K_F1:
                trick1()
            if event.key == ord('1'):
                orb.set_rgb_led(0, 0, 0)
            if event.key == ord('2'):
                orb.set_rgb_led(128, 100, 92)
            if event.key == ord('3'):
                orb.set_rgb_led(255, 200, 180)
            if event.key == ord('4'):
                orb.set_rgb_led(255, 0, 0)
            if event.key == ord('5'):
                orb.set_rgb_led(0, 255, 0)
            if event.key == ord('6'):
                orb.set_rgb_led(0, 0, 255)

        if event.type == pygame.USEREVENT:
            if stab:
                rn = r / 256
                speed = int(255 * (rn * rn * rn))
                orb.roll(speed, (ang + dang) % 360)
                #orb.set_rgb_led(r, r, r)
            else:
                orb.set_stabilization(False)

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
            if not down:
                r = 0

    ks = pygame.key.get_pressed()
    if ks[pygame.K_LEFT]:
        ang = ang - 3
    if ks[pygame.K_RIGHT]:
        ang = ang + 3
    if ks[pygame.K_UP]:
        r = r + 3
    if ks[pygame.K_DOWN]:
        r = r - 3
    if ks[ord(',')]:
        ang = 0
        r = 0
        dang = dang - 3
    if ks[ord('.')]:
        ang = 0
        r = 0
        dang = dang + 3

    if r < 0:
        r = 0

    if r > 250:
        r = 250

    if r > 0:
        stab = True

    a = (ang2 - 90) / 180 * 3.1415
    x = int(400 + cos(a) * 250)
    y = int(300 + sin(a) * 250)
    pygame.draw.lines(screen, (0, 100, 0), False, [(400, 300), (x, y)], 3)

    a = (ang - 90) / 180 * 3.1415
    x = int(400 + cos(a) * r)
    y = int(300 + sin(a) * r)
    pygame.draw.lines(screen, (0, 0, 255), False, [(400, 300), (x, y)], 10)

    pygame.draw.circle(screen, (255, 0, 0), (x, y), 20, 5)
    pygame.display.update()

    pygame.display.flip()
    pygame.time.wait(10)

# End

