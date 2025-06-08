import { Entity, PrimaryGeneratedColumn, Column } from 'typeorm';
import { ApiProperty } from '@nestjs/swagger';

@Entity()
export class Inventory {
  @ApiProperty({ example: 1, description: 'Unique ID of the inventory item' })
  @PrimaryGeneratedColumn()
  id: number;

  @ApiProperty({ example: 'Product A', description: 'Name of the product' })
  @Column()
  productName: string;

  @ApiProperty({ example: 100, description: 'Quantity in stock' })
  @Column('int')
  quantity: number;
} 